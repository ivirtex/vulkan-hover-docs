# VK_FUCHSIA_buffer_collection(3) Manual Page

## Name

VK_FUCHSIA_buffer_collection - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

367

## <a href="#_revision" class="anchor"></a>Revision

2

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_FUCHSIA_external_memory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_FUCHSIA_external_memory.html)  
and  
    
[VK_KHR_sampler_ycbcr_conversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_sampler_ycbcr_conversion.html)  
     or  
     [Version 1.1](#versions-1.1)  

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_EXT_debug_report

## <a href="#_contact" class="anchor"></a>Contact

- John Rosasco <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_FUCHSIA_buffer_collection%5D%20@rosasco%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_FUCHSIA_buffer_collection%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>rosasco</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-09-23

**IP Status**  
No known IP claims.

**Contributors**  
- Craig Stout, Google

- John Bauman, Google

- John Rosasco, Google

## <a href="#_description" class="anchor"></a>Description

A buffer collection is a collection of one or more buffers which were
allocated together as a group and which all have the same properties.
These properties describe the buffers' internal representation such as
its dimensions and memory layout. This ensures that all of the buffers
can be used interchangeably by tasks that require swapping among
multiple buffers, such as double-buffered graphics rendering.

By sharing such a collection of buffers between components,
communication about buffer lifecycle can be made much simpler and more
efficient. For example, when a content producer finishes writing to a
buffer, it can message the consumer of the buffer with the buffer index,
rather than passing a handle to the shared memory.

On Fuchsia, the Sysmem service uses buffer collections as a core
construct in its design. VK_FUCHSIA_buffer_collection is the Vulkan
extension that allows Vulkan applications to interoperate with the
Sysmem service on Fuchsia.

## <a href="#_new_object_types" class="anchor"></a>New Object Types

- [VkBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionFUCHSIA.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCreateBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateBufferCollectionFUCHSIA.html)

- [vkDestroyBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyBufferCollectionFUCHSIA.html)

- [vkGetBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetBufferCollectionPropertiesFUCHSIA.html)

- [vkSetBufferCollectionBufferConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetBufferCollectionBufferConstraintsFUCHSIA.html)

- [vkSetBufferCollectionImageConstraintsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkSetBufferCollectionImageConstraintsFUCHSIA.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkBufferCollectionConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionConstraintsInfoFUCHSIA.html)

- [VkBufferCollectionCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionCreateInfoFUCHSIA.html)

- [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionPropertiesFUCHSIA.html)

- [VkBufferConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferConstraintsInfoFUCHSIA.html)

- [VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageConstraintsInfoFUCHSIA.html)

- [VkImageFormatConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatConstraintsInfoFUCHSIA.html)

- [VkSysmemColorSpaceFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSysmemColorSpaceFUCHSIA.html)

- Extending [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html):

  - [VkBufferCollectionBufferCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionBufferCreateInfoFUCHSIA.html)

- Extending [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html):

  - [VkBufferCollectionImageCreateInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionImageCreateInfoFUCHSIA.html)

- Extending [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html):

  - [VkImportMemoryBufferCollectionFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImportMemoryBufferCollectionFUCHSIA.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkImageConstraintsInfoFlagBitsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageConstraintsInfoFlagBitsFUCHSIA.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkImageConstraintsInfoFlagsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageConstraintsInfoFlagsFUCHSIA.html)

- [VkImageFormatConstraintsFlagsFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatConstraintsFlagsFUCHSIA.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_FUCHSIA_BUFFER_COLLECTION_EXTENSION_NAME`

- `VK_FUCHSIA_BUFFER_COLLECTION_SPEC_VERSION`

- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html):

  - `VK_OBJECT_TYPE_BUFFER_COLLECTION_FUCHSIA`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

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

If [VK_EXT_debug_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_report.html) is supported:

- Extending
  [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html):

  - `VK_DEBUG_REPORT_OBJECT_TYPE_BUFFER_COLLECTION_FUCHSIA_EXT`

## <a href="#_issues" class="anchor"></a>Issues

1\) When configuring a
[VkImageConstraintsInfoFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageConstraintsInfoFUCHSIA.html)
structure for constraint setting, should a NULL `pFormatConstraints`
parameter be allowed ?

**RESOLVED**: No. Specifying a NULL `pFormatConstraints` results in
logical complexity of interpreting the relationship between the
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage` settings of the
elements of the `pImageCreateInfos` array and the implied or desired
[VkFormatFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlags.html).

The explicit requirement for `pFormatConstraints` to be non-NULL
simplifies the implied logic of the implementation and expectations for
the Vulkan application.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 2, 2021-09-23 (John Rosasco)

  - Review passes

- Revision 1, 2021-03-09 (John Rosasco)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_FUCHSIA_buffer_collection"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
