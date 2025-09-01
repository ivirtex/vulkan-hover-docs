# VK\_FUCHSIA\_buffer\_collection(3) Manual Page

## Name

VK\_FUCHSIA\_buffer\_collection - device extension



## [](#_registered_extension_number)Registered Extension Number

367

## [](#_revision)Revision

2

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_FUCHSIA\_external\_memory](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_FUCHSIA_external_memory.html)  
and  
     [VK\_KHR\_sampler\_ycbcr\_conversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_sampler_ycbcr_conversion.html)  
     or  
     [Vulkan Version 1.1](#versions-1.1)

## [](#_api_interactions)API Interactions

- Interacts with VK\_EXT\_debug\_report

## [](#_contact)Contact

- John Rosasco [\[GitHub\]rosasco](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_FUCHSIA_buffer_collection%5D%20%40rosasco%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_FUCHSIA_buffer_collection%20extension%2A)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-09-23

**IP Status**

No known IP claims.

**Contributors**

- Craig Stout, Google
- John Bauman, Google
- John Rosasco, Google

## [](#_description)Description

A buffer collection is a collection of one or more buffers which were allocated together as a group and which all have the same properties. These properties describe the buffers' internal representation such as its dimensions and memory layout. This ensures that all of the buffers can be used interchangeably by tasks that require swapping among multiple buffers, such as double-buffered graphics rendering.

By sharing such a collection of buffers between components, communication about buffer lifecycle can be made much simpler and more efficient. For example, when a content producer finishes writing to a buffer, it can message the consumer of the buffer with the buffer index, rather than passing a handle to the shared memory.

On Fuchsia, the Sysmem service uses buffer collections as a core construct in its design. VK\_FUCHSIA\_buffer\_collection is the Vulkan extension that allows Vulkan applications to interoperate with the Sysmem service on Fuchsia.

## [](#_new_object_types)New Object Types

- [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionFUCHSIA.html)

## [](#_new_commands)New Commands

- [vkCreateBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateBufferCollectionFUCHSIA.html)
- [vkDestroyBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkDestroyBufferCollectionFUCHSIA.html)
- [vkGetBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferCollectionPropertiesFUCHSIA.html)
- [vkSetBufferCollectionBufferConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionBufferConstraintsFUCHSIA.html)
- [vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html)

## [](#_new_structures)New Structures

- [VkBufferCollectionConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionConstraintsInfoFUCHSIA.html)
- [VkBufferCollectionCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionCreateInfoFUCHSIA.html)
- [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionPropertiesFUCHSIA.html)
- [VkBufferConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferConstraintsInfoFUCHSIA.html)
- [VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageConstraintsInfoFUCHSIA.html)
- [VkImageFormatConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatConstraintsInfoFUCHSIA.html)
- [VkSysmemColorSpaceFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSysmemColorSpaceFUCHSIA.html)
- Extending [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html):
  
  - [VkBufferCollectionBufferCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionBufferCreateInfoFUCHSIA.html)
- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html):
  
  - [VkBufferCollectionImageCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionImageCreateInfoFUCHSIA.html)
- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html):
  
  - [VkImportMemoryBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImportMemoryBufferCollectionFUCHSIA.html)

## [](#_new_enums)New Enums

- [VkImageConstraintsInfoFlagBitsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageConstraintsInfoFlagBitsFUCHSIA.html)

## [](#_new_bitmasks)New Bitmasks

- [VkImageConstraintsInfoFlagsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageConstraintsInfoFlagsFUCHSIA.html)
- [VkImageFormatConstraintsFlagsFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageFormatConstraintsFlagsFUCHSIA.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_FUCHSIA_BUFFER_COLLECTION_EXTENSION_NAME`
- `VK_FUCHSIA_BUFFER_COLLECTION_SPEC_VERSION`
- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkObjectType.html):
  
  - `VK_OBJECT_TYPE_BUFFER_COLLECTION_FUCHSIA`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_BUFFER_COLLECTION_BUFFER_CREATE_INFO_FUCHSIA`
  - `VK_STRUCTURE_TYPE_BUFFER_COLLECTION_CONSTRAINTS_INFO_FUCHSIA`
  - `VK_STRUCTURE_TYPE_BUFFER_COLLECTION_CREATE_INFO_FUCHSIA`
  - `VK_STRUCTURE_TYPE_BUFFER_COLLECTION_IMAGE_CREATE_INFO_FUCHSIA`
  - `VK_STRUCTURE_TYPE_BUFFER_COLLECTION_PROPERTIES_FUCHSIA`
  - `VK_STRUCTURE_TYPE_BUFFER_CONSTRAINTS_INFO_FUCHSIA`
  - `VK_STRUCTURE_TYPE_IMAGE_CONSTRAINTS_INFO_FUCHSIA`
  - `VK_STRUCTURE_TYPE_IMAGE_FORMAT_CONSTRAINTS_INFO_FUCHSIA`
  - `VK_STRUCTURE_TYPE_IMPORT_MEMORY_BUFFER_COLLECTION_FUCHSIA`
  - `VK_STRUCTURE_TYPE_SYSMEM_COLOR_SPACE_FUCHSIA`

If [VK\_EXT\_debug\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_debug_report.html) is supported:

- Extending [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDebugReportObjectTypeEXT.html):
  
  - `VK_DEBUG_REPORT_OBJECT_TYPE_BUFFER_COLLECTION_FUCHSIA_EXT`

## [](#_issues)Issues

1\) When configuring a [VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageConstraintsInfoFUCHSIA.html) structure for constraint setting, should a NULL `pFormatConstraints` parameter be allowed ?

**RESOLVED**: No. Specifying a NULL `pFormatConstraints` results in logical complexity of interpreting the relationship between the [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage` settings of the elements of the `pImageCreateInfos` array and the implied or desired [VkFormatFeatureFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlags.html).

The explicit requirement for `pFormatConstraints` to be non-NULL simplifies the implied logic of the implementation and expectations for the Vulkan application.

## [](#_version_history)Version History

- Revision 2, 2021-09-23 (John Rosasco)
  
  - Review passes
- Revision 1, 2021-03-09 (John Rosasco)
  
  - Initial revision

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_FUCHSIA_buffer_collection).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0