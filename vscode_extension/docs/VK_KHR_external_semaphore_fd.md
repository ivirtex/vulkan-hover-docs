# VK\_KHR\_external\_semaphore\_fd(3) Manual Page

## Name

VK\_KHR\_external\_semaphore\_fd - device extension



## [](#_registered_extension_number)Registered Extension Number

80

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_external\_semaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_semaphore.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- James Jones [\[GitHub\]cubanismo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_semaphore_fd%5D%20%40cubanismo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_semaphore_fd%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-10-21

**IP Status**

No known IP claims.

**Contributors**

- Jesse Hall, Google
- James Jones, NVIDIA
- Jeff Juliano, NVIDIA
- Carsten Rohde, NVIDIA

## [](#_description)Description

An application using external memory may wish to synchronize access to that memory using semaphores. This extension enables an application to export semaphore payload to and import semaphore payload from POSIX file descriptors.

## [](#_new_commands)New Commands

- [vkGetSemaphoreFdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSemaphoreFdKHR.html)
- [vkImportSemaphoreFdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkImportSemaphoreFdKHR.html)

## [](#_new_structures)New Structures

- [VkImportSemaphoreFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportSemaphoreFdInfoKHR.html)
- [VkSemaphoreGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphoreGetFdInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_EXTERNAL_SEMAPHORE_FD_EXTENSION_NAME`
- `VK_KHR_EXTERNAL_SEMAPHORE_FD_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_IMPORT_SEMAPHORE_FD_INFO_KHR`
  - `VK_STRUCTURE_TYPE_SEMAPHORE_GET_FD_INFO_KHR`

## [](#_issues)Issues

1\) Does the application need to close the file descriptor returned by [vkGetSemaphoreFdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetSemaphoreFdKHR.html)?

**RESOLVED**: Yes, unless it is passed back in to a driver instance to import the semaphore. A successful get call transfers ownership of the file descriptor to the application, and a successful import transfers it back to the driver. Destroying the original semaphore object will not close the file descriptor or remove its reference to the underlying semaphore resource associated with it.

## [](#_version_history)Version History

- Revision 1, 2016-10-21 (Jesse Hall)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_external_semaphore_fd).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0