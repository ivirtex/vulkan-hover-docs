# VK_KHR_external_memory_fd(3) Manual Page

## Name

VK_KHR_external_memory_fd - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

75

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
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_memory_fd%5D%20@cubanismo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_memory_fd%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cubanismo</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-10-21

**IP Status**  
No known IP claims.

**Contributors**  
- James Jones, NVIDIA

- Jeff Juliano, NVIDIA

## <a href="#_description" class="anchor"></a>Description

An application may wish to reference device memory in multiple Vulkan
logical devices or instances, in multiple processes, and/or in multiple
APIs. This extension enables an application to export POSIX file
descriptor handles from Vulkan memory objects and to import Vulkan
memory objects from POSIX file descriptor handles exported from other
Vulkan memory objects or from similar resources in other APIs.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkGetMemoryFdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryFdKHR.html)

- [vkGetMemoryFdPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryFdPropertiesKHR.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkMemoryFdPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryFdPropertiesKHR.html)

- [VkMemoryGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryGetFdInfoKHR.html)

- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html):

  - [VkImportMemoryFdInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryFdInfoKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_EXTERNAL_MEMORY_FD_EXTENSION_NAME`

- `VK_KHR_EXTERNAL_MEMORY_FD_SPEC_VERSION`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_IMPORT_MEMORY_FD_INFO_KHR`

  - `VK_STRUCTURE_TYPE_MEMORY_FD_PROPERTIES_KHR`

  - `VK_STRUCTURE_TYPE_MEMORY_GET_FD_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) Does the application need to close the file descriptor returned by
[vkGetMemoryFdKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMemoryFdKHR.html)?

**RESOLVED**: Yes, unless it is passed back in to a driver instance to
import the memory. A successful get call transfers ownership of the file
descriptor to the application, and a successful import transfers it back
to the driver. Destroying the original memory object will not close the
file descriptor or remove its reference to the underlying memory
resource associated with it.

2\) Do drivers ever need to expose multiple file descriptors per memory
object?

**RESOLVED**: No. This would indicate there are actually multiple memory
objects, rather than a single memory object.

3\) How should the valid size and memory type for POSIX file descriptor
memory handles created outside of Vulkan be specified?

**RESOLVED**: The valid memory types are queried directly from the
external handle. The size will be specified by future extensions that
introduce such external memory handle types.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-10-21 (James Jones)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_external_memory_fd"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
