include (InstallRequiredSystemLibraries)

########################### General Settings ###########################
set(CPACK_PACKAGE_NAME "MILAGRO")
set(CPACK_PACKAGE_VERSION "${PROJECT_VERSION}")
set(CPACK_PACKAGE_RELEASE 1)
set(CPACK_DESCRIPTION_SUMMARY "${CMAKE_CURRENT_SOURCE_DIR}/README.md")
set(CPACK_RESOURCE_FILE_LICENSE "${CMAKE_CURRENT_SOURCE_DIR}/LICENSE")
set(CPACK_PACKAGE_VENDOR "MILAGRO")
set(CPACK_PACKAGE_CONTACT "dev@milagro.apache.org")
set(CPACK_SYSTEM_NAME "${CMAKE_SYSTEM_NAME}")

if (BUILD_PYTHON)
  set(CPACK_RPM_PACKAGE_REQUIRES "python >= 2.7.0")
endif (BUILD_PYTHON)

set(CPACK_PACKAGE_FILE_NAME "${CPACK_PACKAGE_NAME}-${CPACK_PACKAGE_VERSION}-${CPACK_PACKAGE_RELEASE}.${CMAKE_SYSTEM_PROCESSOR}")

########################### Linux Settings ###########################
if(${CMAKE_SYSTEM_NAME} MATCHES "Linux")
  set(CPACK_PACKAGING_INSTALL_PREFIX ${CMAKE_INSTALL_PREFIX})

  # Prevents CPack from generating file conflicts
  set(CPACK_RPM_EXCLUDE_FROM_AUTO_FILELIST_ADDITION "${CPACK_PACKAGING_INSTALL_PREFIX}")
  list(APPEND CPACK_RPM_EXCLUDE_FROM_AUTO_FILELIST_ADDITION "${CPACK_PACKAGING_INSTALL_PREFIX}/bin")
  list(APPEND CPACK_RPM_EXCLUDE_FROM_AUTO_FILELIST_ADDITION "${CPACK_PACKAGING_INSTALL_PREFIX}/include")
  list(APPEND CPACK_RPM_EXCLUDE_FROM_AUTO_FILELIST_ADDITION "${CPACK_PACKAGING_INSTALL_PREFIX}/lib")
  list(APPEND CPACK_RPM_EXCLUDE_FROM_AUTO_FILELIST_ADDITION "${PYTHON_SITE_LIB}")
  list(APPEND CPACK_RPM_EXCLUDE_FROM_AUTO_FILELIST_ADDITION "${PYTHON_SITE_PACKAGES}")
  set(CPACK_GENERATOR "RPM")
endif()

include (CPack)
