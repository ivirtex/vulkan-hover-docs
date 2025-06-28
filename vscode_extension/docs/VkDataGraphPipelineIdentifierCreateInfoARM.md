# VkDataGraphPipelineIdentifierCreateInfoARM(3) Manual Page

## Name

VkDataGraphPipelineIdentifierCreateInfoARM - Structure specifying an identifier for the newly created data graph pipeline



## [](#_c_specification)C Specification

The `VkDataGraphPipelineIdentifierCreateInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkDataGraphPipelineIdentifierCreateInfoARM {
    VkStructureType    sType;
    const void*        pNext;
    uint32_t           identifierSize;
    const uint8_t*     pIdentifier;
} VkDataGraphPipelineIdentifierCreateInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `identifierSize` is the size in bytes of the identifier data accessible via `pIdentifier`.
- `pIdentifer` is a pointer to `identifierSize` bytes of data that describe the pipeline being created.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDataGraphPipelineIdentifierCreateInfoARM-pIdentifer-09877)VUID-VkDataGraphPipelineIdentifierCreateInfoARM-pIdentifer-09877  
  The data provided via `pIdentifer` **must** have been obtained by calling [vkGetDataGraphPipelinePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDataGraphPipelinePropertiesARM.html) to query the value of the `VK_DATA_GRAPH_PIPELINE_PROPERTY_IDENTIFIER_ARM` property

Valid Usage (Implicit)

- [](#VUID-VkDataGraphPipelineIdentifierCreateInfoARM-sType-sType)VUID-VkDataGraphPipelineIdentifierCreateInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_IDENTIFIER_CREATE_INFO_ARM`
- [](#VUID-VkDataGraphPipelineIdentifierCreateInfoARM-pIdentifier-parameter)VUID-VkDataGraphPipelineIdentifierCreateInfoARM-pIdentifier-parameter  
  `pIdentifier` **must** be a valid pointer to an array of `identifierSize` `uint8_t` values
- [](#VUID-VkDataGraphPipelineIdentifierCreateInfoARM-identifierSize-arraylength)VUID-VkDataGraphPipelineIdentifierCreateInfoARM-identifierSize-arraylength  
  `identifierSize` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphPipelineIdentifierCreateInfoARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0