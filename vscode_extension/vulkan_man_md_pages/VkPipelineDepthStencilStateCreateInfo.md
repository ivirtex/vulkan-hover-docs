# VkPipelineDepthStencilStateCreateInfo(3) Manual Page

## Name

VkPipelineDepthStencilStateCreateInfo - Structure specifying parameters
of a newly created pipeline depth stencil state



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineDepthStencilStateCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkPipelineDepthStencilStateCreateInfo {
    VkStructureType                           sType;
    const void*                               pNext;
    VkPipelineDepthStencilStateCreateFlags    flags;
    VkBool32                                  depthTestEnable;
    VkBool32                                  depthWriteEnable;
    VkCompareOp                               depthCompareOp;
    VkBool32                                  depthBoundsTestEnable;
    VkBool32                                  stencilTestEnable;
    VkStencilOpState                          front;
    VkStencilOpState                          back;
    float                                     minDepthBounds;
    float                                     maxDepthBounds;
} VkPipelineDepthStencilStateCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is a bitmask of
  [VkPipelineDepthStencilStateCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateFlagBits.html)
  specifying additional depth/stencil state information.

- `depthTestEnable` controls whether <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-depth"
  target="_blank" rel="noopener">depth testing</a> is enabled.

- `depthWriteEnable` controls whether <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-depth-write"
  target="_blank" rel="noopener">depth writes</a> are enabled when
  `depthTestEnable` is `VK_TRUE`. Depth writes are always disabled when
  `depthTestEnable` is `VK_FALSE`.

- `depthCompareOp` is a [VkCompareOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html) value specifying
  the comparison operator to use in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-depth-comparison"
  target="_blank" rel="noopener">Depth Comparison</a> step of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-depth"
  target="_blank" rel="noopener">depth test</a>.

- `depthBoundsTestEnable` controls whether <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-dbt"
  target="_blank" rel="noopener">depth bounds testing</a> is enabled.

- `stencilTestEnable` controls whether <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-stencil"
  target="_blank" rel="noopener">stencil testing</a> is enabled.

- `front` and `back` are [VkStencilOpState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOpState.html)
  values controlling the corresponding parameters of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-stencil"
  target="_blank" rel="noopener">stencil test</a>.

- `minDepthBounds` is the minimum depth bound used in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-dbt"
  target="_blank" rel="noopener">depth bounds test</a>.

- `maxDepthBounds` is the maximum depth bound used in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-dbt"
  target="_blank" rel="noopener">depth bounds test</a>.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkPipelineDepthStencilStateCreateInfo-depthBoundsTestEnable-00598"
  id="VUID-VkPipelineDepthStencilStateCreateInfo-depthBoundsTestEnable-00598"></a>
  VUID-VkPipelineDepthStencilStateCreateInfo-depthBoundsTestEnable-00598  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-depthBounds"
  target="_blank" rel="noopener"><code>depthBounds</code></a> feature is
  not enabled, `depthBoundsTestEnable` **must** be `VK_FALSE`

- <a
  href="#VUID-VkPipelineDepthStencilStateCreateInfo-separateStencilMaskRef-04453"
  id="VUID-VkPipelineDepthStencilStateCreateInfo-separateStencilMaskRef-04453"></a>
  VUID-VkPipelineDepthStencilStateCreateInfo-separateStencilMaskRef-04453  
  If the [`VK_KHR_portability_subset`](VK_KHR_portability_subset.html)
  extension is enabled, and
  [VkPhysicalDevicePortabilitySubsetFeaturesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevicePortabilitySubsetFeaturesKHR.html)::`separateStencilMaskRef`
  is `VK_FALSE`, and the value of
  [VkPipelineDepthStencilStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateInfo.html)::`stencilTestEnable`
  is `VK_TRUE`, and the value of
  [VkPipelineRasterizationStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineRasterizationStateCreateInfo.html)::`cullMode`
  is `VK_CULL_MODE_NONE`, the value of `reference` in each of the
  [VkStencilOpState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOpState.html) structs in `front` and
  `back` **must** be the same

- <a
  href="#VUID-VkPipelineDepthStencilStateCreateInfo-rasterizationOrderDepthAttachmentAccess-06463"
  id="VUID-VkPipelineDepthStencilStateCreateInfo-rasterizationOrderDepthAttachmentAccess-06463"></a>
  VUID-VkPipelineDepthStencilStateCreateInfo-rasterizationOrderDepthAttachmentAccess-06463  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-rasterizationOrderDepthAttachmentAccess"
  target="_blank"
  rel="noopener"><code>rasterizationOrderDepthAttachmentAccess</code></a>
  feature is not enabled, `flags` **must** not include
  `VK_PIPELINE_DEPTH_STENCIL_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_DEPTH_ACCESS_BIT_EXT`

- <a
  href="#VUID-VkPipelineDepthStencilStateCreateInfo-rasterizationOrderStencilAttachmentAccess-06464"
  id="VUID-VkPipelineDepthStencilStateCreateInfo-rasterizationOrderStencilAttachmentAccess-06464"></a>
  VUID-VkPipelineDepthStencilStateCreateInfo-rasterizationOrderStencilAttachmentAccess-06464  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-rasterizationOrderStencilAttachmentAccess"
  target="_blank"
  rel="noopener"><code>rasterizationOrderStencilAttachmentAccess</code></a>
  feature is not enabled, `flags` **must** not include
  `VK_PIPELINE_DEPTH_STENCIL_STATE_CREATE_RASTERIZATION_ORDER_ATTACHMENT_STENCIL_ACCESS_BIT_EXT`

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineDepthStencilStateCreateInfo-sType-sType"
  id="VUID-VkPipelineDepthStencilStateCreateInfo-sType-sType"></a>
  VUID-VkPipelineDepthStencilStateCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_DEPTH_STENCIL_STATE_CREATE_INFO`

- <a href="#VUID-VkPipelineDepthStencilStateCreateInfo-pNext-pNext"
  id="VUID-VkPipelineDepthStencilStateCreateInfo-pNext-pNext"></a>
  VUID-VkPipelineDepthStencilStateCreateInfo-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkPipelineDepthStencilStateCreateInfo-flags-parameter"
  id="VUID-VkPipelineDepthStencilStateCreateInfo-flags-parameter"></a>
  VUID-VkPipelineDepthStencilStateCreateInfo-flags-parameter  
  `flags` **must** be a valid combination of
  [VkPipelineDepthStencilStateCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateFlagBits.html)
  values

- <a
  href="#VUID-VkPipelineDepthStencilStateCreateInfo-depthCompareOp-parameter"
  id="VUID-VkPipelineDepthStencilStateCreateInfo-depthCompareOp-parameter"></a>
  VUID-VkPipelineDepthStencilStateCreateInfo-depthCompareOp-parameter  
  `depthCompareOp` **must** be a valid [VkCompareOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html)
  value

- <a href="#VUID-VkPipelineDepthStencilStateCreateInfo-front-parameter"
  id="VUID-VkPipelineDepthStencilStateCreateInfo-front-parameter"></a>
  VUID-VkPipelineDepthStencilStateCreateInfo-front-parameter  
  `front` **must** be a valid [VkStencilOpState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOpState.html)
  structure

- <a href="#VUID-VkPipelineDepthStencilStateCreateInfo-back-parameter"
  id="VUID-VkPipelineDepthStencilStateCreateInfo-back-parameter"></a>
  VUID-VkPipelineDepthStencilStateCreateInfo-back-parameter  
  `back` **must** be a valid [VkStencilOpState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOpState.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkCompareOp](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompareOp.html),
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
[VkPipelineDepthStencilStateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDepthStencilStateCreateFlags.html),
[VkStencilOpState](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStencilOpState.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineDepthStencilStateCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
