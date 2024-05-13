# VkPipelineShaderStageRequiredSubgroupSizeCreateInfo(3) Manual Page

## Name

VkPipelineShaderStageRequiredSubgroupSizeCreateInfo - Structure
specifying the required subgroup size of a newly created pipeline shader
stage



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineShaderStageRequiredSubgroupSizeCreateInfo` structure is
defined as:

``` c
// Provided by VK_VERSION_1_3
typedef struct VkPipelineShaderStageRequiredSubgroupSizeCreateInfo {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           requiredSubgroupSize;
} VkPipelineShaderStageRequiredSubgroupSizeCreateInfo;
```

or the equivalent

``` c
// Provided by VK_EXT_subgroup_size_control
typedef VkPipelineShaderStageRequiredSubgroupSizeCreateInfo VkPipelineShaderStageRequiredSubgroupSizeCreateInfoEXT;
```

or the equiavlent

``` c
// Provided by VK_EXT_shader_object
typedef VkPipelineShaderStageRequiredSubgroupSizeCreateInfo VkShaderRequiredSubgroupSizeCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="pipelines-required-subgroup-size"></span>
  `requiredSubgroupSize` is an unsigned integer value specifying the
  required subgroup size for the newly created pipeline shader stage.

## <a href="#_description" class="anchor"></a>Description

If a `VkPipelineShaderStageRequiredSubgroupSizeCreateInfo` structure is
included in the `pNext` chain of
[VkPipelineShaderStageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineShaderStageCreateInfo.html),
it specifies that the pipeline shader stage being compiled has a
required subgroup size.

If a `VkShaderRequiredSubgroupSizeCreateInfoEXT` structure is included
in the `pNext` chain of
[VkShaderCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShaderCreateInfoEXT.html), it specifies that
the shader being compiled has a required subgroup size.

Valid Usage

- <a
  href="#VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-requiredSubgroupSize-02760"
  id="VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-requiredSubgroupSize-02760"></a>
  VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-requiredSubgroupSize-02760  
  `requiredSubgroupSize` **must** be a power-of-two integer

- <a
  href="#VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-requiredSubgroupSize-02761"
  id="VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-requiredSubgroupSize-02761"></a>
  VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-requiredSubgroupSize-02761  
  `requiredSubgroupSize` **must** be greater or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-minSubgroupSize"
  target="_blank" rel="noopener"><code>minSubgroupSize</code></a>

- <a
  href="#VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-requiredSubgroupSize-02762"
  id="VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-requiredSubgroupSize-02762"></a>
  VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-requiredSubgroupSize-02762  
  `requiredSubgroupSize` **must** be less than or equal to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#limits-maxSubgroupSize"
  target="_blank" rel="noopener"><code>maxSubgroupSize</code></a>

Valid Usage (Implicit)

- <a
  href="#VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-sType-sType"
  id="VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-sType-sType"></a>
  VUID-VkPipelineShaderStageRequiredSubgroupSizeCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_SHADER_STAGE_REQUIRED_SUBGROUP_SIZE_CREATE_INFO`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_subgroup_size_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_subgroup_size_control.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineShaderStageRequiredSubgroupSizeCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
