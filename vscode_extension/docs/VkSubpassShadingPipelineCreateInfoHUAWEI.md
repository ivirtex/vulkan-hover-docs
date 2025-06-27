# VkSubpassShadingPipelineCreateInfoHUAWEI(3) Manual Page

## Name

VkSubpassShadingPipelineCreateInfoHUAWEI - Structure specifying
parameters of a newly created subpass shading pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

A subpass shading pipeline is a compute pipeline which **must** be
called only in a subpass of a render pass with work dimensions specified
by render area size. The subpass shading pipeline shader is a compute
shader allowed to access input attachments specified in the calling
subpass. To create a subpass shading pipeline, call
[vkCreateComputePipelines](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateComputePipelines.html) with
[VkSubpassShadingPipelineCreateInfoHUAWEI](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassShadingPipelineCreateInfoHUAWEI.html)
in the `pNext` chain of
[VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html).

The `VkSubpassShadingPipelineCreateInfoHUAWEI` structure is defined as:

``` c
// Provided by VK_HUAWEI_subpass_shading
typedef struct VkSubpassShadingPipelineCreateInfoHUAWEI {
    VkStructureType    sType;
    void*              pNext;
    VkRenderPass       renderPass;
    uint32_t           subpass;
} VkSubpassShadingPipelineCreateInfoHUAWEI;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `renderPass` is a handle to a render pass object describing the
  environment in which the pipeline will be used. The pipeline **must**
  only be used with a render pass instance compatible with the one
  provided. See <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-compatibility"
  target="_blank" rel="noopener">Render Pass Compatibility</a> for more
  information.

- `subpass` is the index of the subpass in the render pass where this
  pipeline will be used.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkSubpassShadingPipelineCreateInfoHUAWEI-subpass-04946"
  id="VUID-VkSubpassShadingPipelineCreateInfoHUAWEI-subpass-04946"></a>
  VUID-VkSubpassShadingPipelineCreateInfoHUAWEI-subpass-04946  
  `subpass` **must** be created with
  `VK_PIPELINE_BIND_POINT_SUBPASS_SHADING_HUAWEI` bind point

Valid Usage (Implicit)

- <a href="#VUID-VkSubpassShadingPipelineCreateInfoHUAWEI-sType-sType"
  id="VUID-VkSubpassShadingPipelineCreateInfoHUAWEI-sType-sType"></a>
  VUID-VkSubpassShadingPipelineCreateInfoHUAWEI-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SUBPASS_SHADING_PIPELINE_CREATE_INFO_HUAWEI`

- <a
  href="#VUID-VkSubpassShadingPipelineCreateInfoHUAWEI-renderPass-parameter"
  id="VUID-VkSubpassShadingPipelineCreateInfoHUAWEI-renderPass-parameter"></a>
  VUID-VkSubpassShadingPipelineCreateInfoHUAWEI-renderPass-parameter  
  `renderPass` **must** be a valid [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html)
  handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_HUAWEI_subpass_shading](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_HUAWEI_subpass_shading.html),
[VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSubpassShadingPipelineCreateInfoHUAWEI"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
