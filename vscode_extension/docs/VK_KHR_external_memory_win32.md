# VK\_KHR\_external\_memory\_win32(3) Manual Page

## Name

VK\_KHR\_external\_memory\_win32 - device extension



## [](#_registered_extension_number)Registered Extension Number

74

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- James Jones [\[GitHub\]cubanismo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_memory_win32%5D%20%40cubanismo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_memory_win32%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-10-21

**IP Status**

No known IP claims.

**Contributors**

- James Jones, NVIDIA
- Jeff Juliano, NVIDIA
- Carsten Rohde, NVIDIA

## [](#_description)Description

An application may wish to reference device memory in multiple Vulkan logical devices or instances, in multiple processes, and/or in multiple APIs. This extension enables an application to export Windows handles from Vulkan memory objects and to import Vulkan memory objects from Windows handles exported from other Vulkan memory objects or from similar resources in other APIs.

## [](#_new_commands)New Commands

- [vkGetMemoryWin32HandleKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryWin32HandleKHR.html)
- [vkGetMemoryWin32HandlePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryWin32HandlePropertiesKHR.html)

## [](#_new_structures)New Structures

- [VkMemoryGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetWin32HandleInfoKHR.html)
- [VkMemoryWin32HandlePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryWin32HandlePropertiesKHR.html)
- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html):
  
  - [VkExportMemoryWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMemoryWin32HandleInfoKHR.html)
  - [VkImportMemoryWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMemoryWin32HandleInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_EXTERNAL_MEMORY_WIN32_EXTENSION_NAME`
- `VK_KHR_EXTERNAL_MEMORY_WIN32_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_MEMORY_GET_WIN32_HANDLE_INFO_KHR`
  - `VK_STRUCTURE_TYPE_MEMORY_WIN32_HANDLE_PROPERTIES_KHR`

## [](#_issues)Issues

1\) Do applications need to call `CloseHandle`() on the values returned from [vkGetMemoryWin32HandleKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryWin32HandleKHR.html) when `handleType` is `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR`?

**RESOLVED**: Yes. A successful get call transfers ownership of the handle to the application. Destroying the memory object will not destroy the handle or the handle’s reference to the underlying memory resource. Unlike file descriptor opaque handles, win32 opaque handle ownership can not be transferred back to a driver by an import operation.

2\) Should the language regarding KMT/Windows 7 handles be moved to a separate extension so that it can be deprecated over time?

**RESOLVED**: No. Support for them can be deprecated by drivers if they choose, by no longer returning them in the supported handle types of the instance level queries.

3\) How should the valid size and memory type for windows memory handles created outside of Vulkan be specified?

**RESOLVED**: The valid memory types are queried directly from the external handle. The size is determined by the associated image or buffer memory requirements for external handle types that require dedicated allocations, and by the size specified when creating the object from which the handle was exported for other external handle types.

## [](#_version_history)Version History

- Revision 1, 2016-10-21 (James Jones)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_external_memory_win32)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0