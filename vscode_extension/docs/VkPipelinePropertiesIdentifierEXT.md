# VkPipelinePropertiesIdentifierEXT(3) Manual Page

## Name

VkPipelinePropertiesIdentifierEXT - Structure used to retrieve pipeline properties



## [](#_c_specification)C Specification

The `VkPipelinePropertiesIdentifierEXT` structure is defined as:

```c++
// Provided by VK_EXT_pipeline_properties
typedef struct VkPipelinePropertiesIdentifierEXT {
    VkStructureType    sType;
    void*              pNext;
    uint8_t            pipelineIdentifier[VK_UUID_SIZE];
} VkPipelinePropertiesIdentifierEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pipelineIdentifier` is an array of `VK_UUID_SIZE` `uint8_t` values into which the pipeline identifier will be written.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkPipelinePropertiesIdentifierEXT-sType-sType)VUID-VkPipelinePropertiesIdentifierEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_PROPERTIES_IDENTIFIER_EXT`
- [](#VUID-VkPipelinePropertiesIdentifierEXT-pNext-pNext)VUID-VkPipelinePropertiesIdentifierEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_EXT\_pipeline\_properties](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_pipeline_properties.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelinePropertiesIdentifierEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0