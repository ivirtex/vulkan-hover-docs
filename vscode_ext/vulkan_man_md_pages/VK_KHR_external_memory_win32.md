# VK_KHR_external_memory_win32(3) Manual Page

## Name

VK_KHR_external_memory_win32 - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

74

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_memory.html)  
or  
[Version 1.1](#versions-1.1)  

## <a href="#_contact" class="anchor"></a>Contact

- James Jones <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_memory_win32%5D%20@cubanismo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_memory_win32%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cubanismo</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-10-21

**IP Status**  
No known IP claims.

**Contributors**  
- James Jones, NVIDIA

- Jeff Juliano, NVIDIA

- Carsten Rohde, NVIDIA

## <a href="#_description" class="anchor"></a>Description

An application may wish to reference device memory in multiple Vulkan
logical devices or instances, in multiple processes, and/or in multiple
APIs. This extension enables an application to export Windows handles
from Vulkan memory objects and to import Vulkan memory objects from
Windows handles exported from other Vulkan memory objects or from
similar resources in other APIs.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetMemoryWin32HandleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryWin32HandleKHR.html)

- [vkGetMemoryWin32HandlePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryWin32HandlePropertiesKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkMemoryGetWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetWin32HandleInfoKHR.html)

- [VkMemoryWin32HandlePropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryWin32HandlePropertiesKHR.html)

- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html):

  - [VkExportMemoryWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMemoryWin32HandleInfoKHR.html)

  - [VkImportMemoryWin32HandleInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryWin32HandleInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_EXTERNAL_MEMORY_WIN32_EXTENSION_NAME`

- `VK_KHR_EXTERNAL_MEMORY_WIN32_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_MEMORY_GET_WIN32_HANDLE_INFO_KHR`

  - `VK_STRUCTURE_TYPE_MEMORY_WIN32_HANDLE_PROPERTIES_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) Do applications need to call `CloseHandle`() on the values returned
from [vkGetMemoryWin32HandleKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryWin32HandleKHR.html) when
`handleType` is `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT_KHR`?

**RESOLVED**: Yes. A successful get call transfers ownership of the
handle to the application. Destroying the memory object will not destroy
the handle or the handle’s reference to the underlying memory resource.
Unlike file descriptor opaque handles, win32 opaque handle ownership can
not be transferred back to a driver by an import operation.

2\) Should the language regarding KMT/Windows 7 handles be moved to a
separate extension so that it can be deprecated over time?

**RESOLVED**: No. Support for them can be deprecated by drivers if they
choose, by no longer returning them in the supported handle types of the
instance level queries.

3\) How should the valid size and memory type for windows memory handles
created outside of Vulkan be specified?

**RESOLVED**: The valid memory types are queried directly from the
external handle. The size is determined by the associated image or
buffer memory requirements for external handle types that require
dedicated allocations, and by the size specified when creating the
object from which the handle was exported for other external handle
types.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-10-21 (James Jones)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_external_memory_win32"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
