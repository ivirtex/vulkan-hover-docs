# VK\_KHR\_external\_fence\_win32(3) Manual Page

## Name

VK\_KHR\_external\_fence\_win32 - device extension



## [](#_registered_extension_number)Registered Extension Number

115

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_external\_fence](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_fence.html)

## [](#_contact)Contact

- Jesse Hall [\[GitHub\]critsec](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_fence_win32%5D%20%40critsec%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_fence_win32%20extension%2A)

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
- Contributors to `VK_KHR_external_semaphore_win32`

## [](#_description)Description

An application using external memory may wish to synchronize access to that memory using fences. This extension enables an application to export fence payload to and import fence payload from Windows handles.

## [](#_new_commands)New Commands

- [vkGetFenceWin32HandleKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetFenceWin32HandleKHR.html)
- [vkImportFenceWin32HandleKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkImportFenceWin32HandleKHR.html)

## [](#_new_structures)New Structures

- [VkFenceGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceGetWin32HandleInfoKHR.html)
- [VkImportFenceWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportFenceWin32HandleInfoKHR.html)
- Extending [VkFenceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFenceCreateInfo.html):
  
  - [VkExportFenceWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportFenceWin32HandleInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_EXTERNAL_FENCE_WIN32_EXTENSION_NAME`
- `VK_KHR_EXTERNAL_FENCE_WIN32_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_EXPORT_FENCE_WIN32_HANDLE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_FENCE_GET_WIN32_HANDLE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_IMPORT_FENCE_WIN32_HANDLE_INFO_KHR`

## [](#_issues)Issues

This extension borrows concepts, semantics, and language from `VK_KHR_external_semaphore_win32`. That extension’s issues apply equally to this extension.

1\) Should D3D12 fence handle types be supported, like they are for semaphores?

**RESOLVED**: No. Doing so would require extending the fence signal and wait operations to provide values to signal / wait for, like `VkD3D12FenceSubmitInfoKHR` does. A D3D12 fence can be signaled by importing it into a [VkSemaphore](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSemaphore.html) instead of a [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html), and applications can check status or wait on the D3D12 fence using non-Vulkan APIs. The convenience of being able to do these operations on `VkFence` objects does not justify the extra API complexity.

## [](#_version_history)Version History

- Revision 1, 2017-05-08 (Jesse Hall)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_external_fence_win32).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0