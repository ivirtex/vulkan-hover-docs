# VkPipelineShaderStageRequiredSubgroupSizeCreateInfo(3) Manual Page

## Name

VkPipelineShaderStageRequiredSubgroupSizeCreateInfo - Structure specifying the required subgroup size of a newly created pipeline shader stage



## [](#_c_specification)C Specification

The `VkPipelineShaderStageRequiredSubgroupSizeCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_3
typedef struct VkPipelineShaderStageRequiredSubgroupSizeCreateInfo {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           requiredSubgroupSize;
} VkPipelineShaderStageRequiredSubgroupSizeCreateInfo;
```

or the equivalent

```c++
// Provided by VK_EXT_subgroup_size_control
typedef VkPipelineShaderStageRequiredSubgroupSizeCreateInfo VkPipelineShaderStageRequiredSubgroupSizeCreateInfoEXT;
```

or the equiavlent

```c++
// Provided by VK_EXT_shader_object
typedef VkPipelineShaderStageRequiredSubgroupSizeCreateInfo VkShaderRequiredSubgroupSizeCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`requiredSubgroupSize` is an unsigned integer value specifying the required subgroup size for the newly created pipeline shader stage.

## [](#_description)Description

If a `VkPipelineShaderStageRequiredSubgroupSizeCreateInfo` structure is included in the `pNext` chain of [VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineShaderStageCreateInfo.html), it specifies that the pipeline shader stage being compiled has a required subgroup size.

If a `VkShaderRequiredSubgroupSizeCreateInfoEXT` structure is included in the `pNext` chain of [VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkShaderCreateInfoEXT.html), it specifies that the shader being compiled has a required subgroup size.

Valid Usage

- [](#VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-requiredSubgroupSize-02760)VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-requiredSubgroupSize-02760  
  `requiredSubgroupSize` **must** be a power-of-two integer
- [](#VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-requiredSubgroupSize-02761)VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-requiredSubgroupSize-02761  
  `requiredSubgroupSize` **must** be greater or equal to [`minSubgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-minSubgroupSize)
- [](#VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-requiredSubgroupSize-02762)VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-requiredSubgroupSize-02762  
  `requiredSubgroupSize` **must** be less than or equal to [`maxSubgroupSize`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#limits-maxSubgroupSize)

Valid Usage (Implicit)

- [](#VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-sType-sType)VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_REQUIRED_SUBGROUP_SIZE_CREATE_INFO`

## [](#_see_also)See Also

[VK\_EXT\_subgroup\_size\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_subgroup_size_control.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPipelineShaderStageRequiredSubgroupSizeCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0