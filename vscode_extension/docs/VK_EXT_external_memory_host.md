# VK\_EXT\_external\_memory\_host(3) Manual Page

## Name

VK\_EXT\_external\_memory\_host - device extension



## [](#_registered_extension_number)Registered Extension Number

179

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_external_memory.html)  
or  
[Vulkan Version 1.1](#versions-1.1)

## [](#_contact)Contact

- Daniel Rakos [\[GitHub\]drakos-amd](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_EXT_external_memory_host%5D%20%40drakos-amd%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_EXT_external_memory_host%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2017-11-10

**IP Status**

No known IP claims.

**Contributors**

- Jaakko Konttinen, AMD
- David Mao, AMD
- Daniel Rakos, AMD
- Tobias Hector, Imagination Technologies
- Faith Ekstrand, Intel
- James Jones, NVIDIA

## [](#_description)Description

This extension enables an application to import host allocations and host mapped foreign device memory to Vulkan memory objects.

## [](#_new_commands)New Commands

- [vkGetMemoryHostPointerPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryHostPointerPropertiesEXT.html)

## [](#_new_structures)New Structures

- [VkMemoryHostPointerPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryHostPointerPropertiesEXT.html)
- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html):
  
  - [VkImportMemoryHostPointerInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMemoryHostPointerInfoEXT.html)
- Extending [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html):
  
  - [VkPhysicalDeviceExternalMemoryHostPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalMemoryHostPropertiesEXT.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_EXT_EXTERNAL_MEMORY_HOST_EXTENSION_NAME`
- `VK_EXT_EXTERNAL_MEMORY_HOST_SPEC_VERSION`
- Extending [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html):
  
  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_ALLOCATION_BIT_EXT`
  - `VK_EXTERNAL_MEMORY_HANDLE_TYPE_HOST_MAPPED_FOREIGN_MEMORY_BIT_EXT`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_IMPORT_MEMORY_HOST_POINTER_INFO_EXT`
  - `VK_STRUCTURE_TYPE_MEMORY_HOST_POINTER_PROPERTIES_EXT`
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_EXTERNAL_MEMORY_HOST_PROPERTIES_EXT`

## [](#_issues)Issues

1\) What memory type has to be used to import host pointers?

**RESOLVED**: Depends on the implementation. Applications have to use the new [vkGetMemoryHostPointerPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMemoryHostPointerPropertiesEXT.html) command to query the supported memory types for a particular host pointer. The reported memory types may include memory types that come from a memory heap that is otherwise not usable for regular memory object allocation and thus such a heap’s size may be zero.

2\) Can the application still access the contents of the host allocation after importing?

**RESOLVED**: Yes. However, usual synchronization requirements apply.

3\) Can the application free the host allocation?

**RESOLVED**: No, it violates valid usage conditions. Using the memory object imported from a host allocation that is already freed thus results in undefined behavior.

4\) Is [vkMapMemory](https://registry.khronos.org/vulkan/specs/latest/man/html/vkMapMemory.html) expected to return the same host address which was specified when importing it to the memory object?

**RESOLVED**: No. Implementations are allowed to return the same address but it is not required. Some implementations might return a different virtual mapping of the allocation, although the same physical pages will be used.

5\) Is there any limitation on the alignment of the host pointer and/or size?

**RESOLVED**: Yes. Both the address and the size have to be an integer multiple of `minImportedHostPointerAlignment`. In addition, some platforms and foreign devices may have additional restrictions.

6\) Can the same host allocation be imported multiple times into a given physical device?

**RESOLVED**: No, at least not guaranteed by this extension. Some platforms do not allow locking the same physical pages for device access multiple times, so attempting to do it may result in undefined behavior.

7\) Does this extension support exporting the new handle type?

**RESOLVED**: No.

8\) Should we include the possibility to import host mapped foreign device memory using this API?

**RESOLVED**: Yes, through a separate handle type. Implementations are still allowed to support only one of the handle types introduced by this extension by not returning import support for a particular handle type as returned in [VkExternalMemoryPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryPropertiesKHR.html).

## [](#_version_history)Version History

- Revision 1, 2017-11-10 (Daniel Rakos)
  
  - Internal revisions

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_EXT_external_memory_host)

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0