# VK\_KHR\_external\_fence\_fd(3) Manual Page

## Name

VK\_KHR\_external\_fence\_fd - device extension



## [](#_registered_extension_number)Registered Extension Number

116

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_external\_fence](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_fence.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Jesse Hall [\[GitHub\]critsec](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_fence_fd%5D%20%40critsec%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_fence_fd%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-05-08

**IP Status**

No known IP claims.

**Contributors**

- Jesse Hall, Google
- James Jones, NVIDIA
- Jeff Juliano, NVIDIA
- Cass Everitt, Oculus
- Contributors to `VK_KHR_external_semaphore_fd`

## [](#_description)Description

An application using external memory may wish to synchronize access to that memory using fences. This extension enables an application to export fence payload to and import fence payload from POSIX file descriptors.

## [](#_new_commands)New Commands

- [vkGetFenceFdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetFenceFdKHR.html)
- [vkImportFenceFdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkImportFenceFdKHR.html)

## [](#_new_structures)New Structures

- [VkFenceGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceGetFdInfoKHR.html)
- [VkImportFenceFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportFenceFdInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_EXTERNAL_FENCE_FD_EXTENSION_NAME`
- `VK_KHR_EXTERNAL_FENCE_FD_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_FENCE_GET_FD_INFO_KHR`
  - `VK_STRUCTURE_TYPE_IMPORT_FENCE_FD_INFO_KHR`

## [](#_issues)Issues

This extension borrows concepts, semantics, and language from `VK_KHR_external_semaphore_fd`. That extension’s issues apply equally to this extension.

## [](#_version_history)Version History

- Revision 1, 2017-05-08 (Jesse Hall)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_external_fence_fd)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0