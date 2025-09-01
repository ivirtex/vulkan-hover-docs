# VkDataGraphPipelinePropertyARM(3) Manual Page

## Name

VkDataGraphPipelinePropertyARM - Enumeration describing the properties of a data graph pipeline that can be queried



## [](#_c_specification)C Specification

Possible values of [VkDataGraphPipelinePropertyQueryResultARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelinePropertyQueryResultARM.html)::`property`, specifying the property of the data graph pipeline being queried, are:

```c++
// Provided by VK_ARM_data_graph
typedef enum VkDataGraphPipelinePropertyARM {
    VK_DATA_GRAPH_PIPELINE_PROPERTY_CREATION_LOG_ARM = 0,
    VK_DATA_GRAPH_PIPELINE_PROPERTY_IDENTIFIER_ARM = 1,
} VkDataGraphPipelinePropertyARM;
```

## [](#_description)Description

- `VK_DATA_GRAPH_PIPELINE_PROPERTY_CREATION_LOG_ARM` corresponds to a human-readable log produced during the creation of a data graph pipeline. It **may** contain information about errors encountered during the creation or other information generally useful for debugging. This property **can** be queried for any data graph pipeline.
- `VK_DATA_GRAPH_PIPELINE_PROPERTY_IDENTIFIER_ARM` corresponds to an opaque identifier for the data graph pipeline. It **can** be used to create a graph pipeline from a pipeline cache without the need to provide any creation data beyond the identifier, using a [VkDataGraphPipelineIdentifierCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelineIdentifierCreateInfoARM.html) structure.

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkDataGraphPipelinePropertyQueryResultARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDataGraphPipelinePropertyQueryResultARM.html), [vkGetDataGraphPipelineAvailablePropertiesARM](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDataGraphPipelineAvailablePropertiesARM.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphPipelinePropertyARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0