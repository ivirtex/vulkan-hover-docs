# VK_HUAWEI_invocation_mask(3) Manual Page

## Name

VK_HUAWEI_invocation_mask - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

371

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html)  
and  
     [VK_KHR_synchronization2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html)  
     or  
     [Version 1.3](#versions-1.3)  

## <a href="#_contact" class="anchor"></a>Contact

- Pan Gao <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_HUAWEI_invocation_mask%5D%20@PanGao-h%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_HUAWEI_invocation_mask%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>PanGao-h</a>

## <a href="#_extension_proposal" class="anchor"></a>Extension Proposal

[VK_HUAWEI_invocation_mask](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_HUAWEI_invocation_mask.adoc)

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-05-27

**Interactions and External Dependencies**  
- This extension requires
  [`VK_KHR_ray_tracing_pipeline`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html),
  which allow to bind an invocation mask image before the ray tracing
  command

- This extension requires
  [`VK_KHR_synchronization2`](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_synchronization2.html), which
  allows new pipeline stage for the invocation mask image

**Contributors**  
- Yunpeng Zhu

- Juntao Li, Huawei

- Liang Chen, Huawei

- Shaozhuang Shi, Huawei

- Hailong Chu, Huawei

## <a href="#_description" class="anchor"></a>Description

The rays to trace may be sparse in some use cases. For example, the
scene only have a few regions to reflect. Providing an invocation mask
image to the ray tracing commands could potentially give the hardware
the hint to do certain optimization without invoking an additional pass
to compact the ray buffer.

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdBindInvocationMaskHUAWEI](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindInvocationMaskHUAWEI.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceFeatures2.html),
  [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceCreateInfo.html):

  - [VkPhysicalDeviceInvocationMaskFeaturesHUAWEI](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceInvocationMaskFeaturesHUAWEI.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_HUAWEI_INVOCATION_MASK_EXTENSION_NAME`

- `VK_HUAWEI_INVOCATION_MASK_SPEC_VERSION`

- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccessFlagBits2.html):

  - `VK_ACCESS_2_INVOCATION_MASK_READ_BIT_HUAWEI`

- Extending [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html):

  - `VK_IMAGE_USAGE_INVOCATION_MASK_BIT_HUAWEI`

- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineStageFlagBits2.html):

  - `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INVOCATION_MASK_FEATURES_HUAWEI`

## <a href="#_examples" class="anchor"></a>Examples

RT mask is updated before each traceRay.

Step 1. Generate InvocationMask.

``` c
//the rt mask image bind as color attachment in the fragment shader
Layout(location = 2) out vec4 outRTmask
vec4 mask = vec4(x,x,x,x);
outRTmask = mask;
```

Step 2. traceRay with InvocationMask

``` c
vkCmdBindPipeline(
    commandBuffers[imageIndex],
    VK_PIPELINE_BIND_POINT_RAY_TRACING_KHR, m_rtPipeline);
    vkCmdBindDescriptorSets(commandBuffers[imageIndex],
    VK_PIPELINE_BIND_POINT_RAY_TRACING_NV,
    m_rtPipelineLayout, 0, 1, &m_rtDescriptorSet,
    0, nullptr);

vkCmdBindInvocationMaskHUAWEI(
    commandBuffers[imageIndex],
    InvocationMaskimageView,
    InvocationMaskimageLayout);
    vkCmdTraceRaysKHR(commandBuffers[imageIndex],
    pRaygenShaderBindingTable,
    pMissShaderBindingTable,
    swapChainExtent.width,
    swapChainExtent.height, 1);
```

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-05-27 (Yunpeng Zhu)

  - Initial draft.

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_HUAWEI_invocation_mask"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
