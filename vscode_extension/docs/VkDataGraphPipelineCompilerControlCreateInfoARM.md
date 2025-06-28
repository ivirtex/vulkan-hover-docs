# VkDataGraphPipelineCompilerControlCreateInfoARM(3) Manual Page

## Name

VkDataGraphPipelineCompilerControlCreateInfoARM - Structure specifying compiler control parameters of a newly created data graph pipeline



## [](#_c_specification)C Specification

The `VkDataGraphPipelineCompilerControlCreateInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_data_graph
typedef struct VkDataGraphPipelineCompilerControlCreateInfoARM {
    VkStructureType    sType;
    const void*        pNext;
    const char*        pVendorOptions;
} VkDataGraphPipelineCompilerControlCreateInfoARM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pVendorOptions` is a null-terminated UTF-8 string specifying implementation-specific options that affect the creation of a data graph pipeline.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDataGraphPipelineCompilerControlCreateInfoARM-sType-sType)VUID-VkDataGraphPipelineCompilerControlCreateInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DATA_GRAPH_PIPELINE_COMPILER_CONTROL_CREATE_INFO_ARM`
- [](#VUID-VkDataGraphPipelineCompilerControlCreateInfoARM-pVendorOptions-parameter)VUID-VkDataGraphPipelineCompilerControlCreateInfoARM-pVendorOptions-parameter  
  `pVendorOptions` **must** be a null-terminated UTF-8 string

## [](#_see_also)See Also

[VK\_ARM\_data\_graph](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_data_graph.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDataGraphPipelineCompilerControlCreateInfoARM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0