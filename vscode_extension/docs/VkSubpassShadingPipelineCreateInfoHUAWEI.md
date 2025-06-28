# VkSubpassShadingPipelineCreateInfoHUAWEI(3) Manual Page

## Name

VkSubpassShadingPipelineCreateInfoHUAWEI - Structure specifying parameters of a newly created subpass shading pipeline



## [](#_c_specification)C Specification

A subpass shading pipeline is a compute pipeline which **must** be called only in a subpass of a render pass with work dimensions specified by render area size. The subpass shading pipeline shader is a compute shader allowed to access input attachments specified in the calling subpass. To create a subpass shading pipeline, call [vkCreateComputePipelines](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateComputePipelines.html) with [VkSubpassShadingPipelineCreateInfoHUAWEI](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassShadingPipelineCreateInfoHUAWEI.html) in the `pNext` chain of [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComputePipelineCreateInfo.html).

The `VkSubpassShadingPipelineCreateInfoHUAWEI` structure is defined as:

```c++
// Provided by VK_HUAWEI_subpass_shading
typedef struct VkSubpassShadingPipelineCreateInfoHUAWEI {
    VkStructureType    sType;
    void*              pNext;
    VkRenderPass       renderPass;
    uint32_t           subpass;
} VkSubpassShadingPipelineCreateInfoHUAWEI;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `renderPass` is a handle to a render pass object describing the environment in which the pipeline will be used. The pipeline **must** only be used with a render pass instance compatible with the one provided. See [Render Pass Compatibility](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-compatibility) for more information. The implementation **must** not access this object outside of the duration of the command this structure is passed to.
- `subpass` is the index of the subpass in the render pass where this pipeline will be used.

## [](#_description)Description

Valid Usage

- [](#VUID-VkSubpassShadingPipelineCreateInfoHUAWEI-subpass-04946)VUID-VkSubpassShadingPipelineCreateInfoHUAWEI-subpass-04946  
  `subpass` **must** be created with `VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI` bind point

Valid Usage (Implicit)

- [](#VUID-VkSubpassShadingPipelineCreateInfoHUAWEI-sType-sType)VUID-VkSubpassShadingPipelineCreateInfoHUAWEI-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SUBPASS_SHADING_PIPELINE_CREATE_INFO_HUAWEI`
- [](#VUID-VkSubpassShadingPipelineCreateInfoHUAWEI-renderPass-parameter)VUID-VkSubpassShadingPipelineCreateInfoHUAWEI-renderPass-parameter  
  `renderPass` **must** be a valid [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) handle

## [](#_see_also)See Also

[VK\_HUAWEI\_subpass\_shading](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_HUAWEI_subpass_shading.html), [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSubpassShadingPipelineCreateInfoHUAWEI)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0