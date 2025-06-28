# VK\_KHR\_external\_memory\_fd(3) Manual Page

## Name

VK\_KHR\_external\_memory\_fd - device extension



## [](#_registered_extension_number)Registered Extension Number

75

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- James Jones [\[GitHub\]cubanismo](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_memory_fd%5D%20%40cubanismo%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_memory_fd%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2016-10-21

**IP Status**

No known IP claims.

**Contributors**

- James Jones, NVIDIA
- Jeff Juliano, NVIDIA

## [](#_description)Description

An application may wish to reference device memory in multiple Vulkan logical devices or instances, in multiple processes, and/or in multiple APIs. This extension enables an application to export POSIX file descriptor handles from Vulkan memory objects and to import Vulkan memory objects from POSIX file descriptor handles exported from other Vulkan memory objects or from similar resources in other APIs.

## [](#_new_commands)New Commands

- [vkGetMemoryFdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryFdKHR.html)
- [vkGetMemoryFdPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryFdPropertiesKHR.html)

## [](#_new_structures)New Structures

- [VkMemoryFdPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryFdPropertiesKHR.html)
- [VkMemoryGetFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryGetFdInfoKHR.html)
- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html):
  
  - [VkImportMemoryFdInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMemoryFdInfoKHR.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_KHR_EXTERNAL_MEMORY_FD_EXTENSION_NAME`
- `VK_KHR_EXTERNAL_MEMORY_FD_SPEC_VERSION`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_IMPORT_MEMORY_FD_INFO_KHR`
  - `VK_STRUCTURE_TYPE_MEMORY_FD_PROPERTIES_KHR`
  - `VK_STRUCTURE_TYPE_MEMORY_GET_FD_INFO_KHR`

## [](#_issues)Issues

1\) Does the application need to close the file descriptor returned by [vkGetMemoryFdKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryFdKHR.html)?

**RESOLVED**: Yes, unless it is passed back in to a driver instance to import the memory. A successful get call transfers ownership of the file descriptor to the application, and a successful import transfers it back to the driver. Destroying the original memory object will not close the file descriptor or remove its reference to the underlying memory resource associated with it.

2\) Do drivers ever need to expose multiple file descriptors per memory object?

**RESOLVED**: No. This would indicate there are actually multiple memory objects, rather than a single memory object.

3\) How should the valid size and memory type for POSIX file descriptor memory handles created outside of Vulkan be specified?

**RESOLVED**: The valid memory types are queried directly from the external handle. The size will be specified by future extensions that introduce such external memory handle types.

## [](#_version_history)Version History

- Revision 1, 2016-10-21 (James Jones)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_KHR_external_memory_fd)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0