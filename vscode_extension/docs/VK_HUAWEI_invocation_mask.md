# VK\_HUAWEI\_invocation\_mask(3) Manual Page

## Name

VK\_HUAWEI\_invocation\_mask - device extension



## [](#_registered_extension_number)Registered Extension Number

371

## [](#_revision)Revision

1

## [](#_ratification_status)Ratification Status

Not ratified

## [](#_extension_and_version_dependencies)Extension and Version Dependencies

[VK\_KHR\_ray\_tracing\_pipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_ray_tracing_pipeline.html)  
and  
     [VK\_KHR\_synchronization2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_synchronization2.html)  
     or  
     [Vulkan Version 1.3](#versions-1.3)

## [](#_contact)Contact

- Pan Gao [\[GitHub\]PanGao-h](https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_HUAWEI_invocation_mask%5D%20%40PanGao-h%0A%2AHere%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_HUAWEI_invocation_mask%20extension%2A)

## [](#_extension_proposal)Extension Proposal

[VK\_HUAWEI\_invocation\_mask](https://github.com/KhronosGroup/Vulkan-Docs/tree/main/proposals/VK_HUAWEI_invocation_mask.adoc)

## [](#_other_extension_metadata)Other Extension Metadata

**Last Modified Date**

2021-05-27

**Interactions and External Dependencies**

- This extension requires `VK_KHR_ray_tracing_pipeline`, which allow to bind an invocation mask image before the ray tracing command
- This extension requires `VK_KHR_synchronization2`, which allows new pipeline stage for the invocation mask image

**Contributors**

- Yunpeng Zhu
- Juntao Li, Huawei
- Liang Chen, Huawei
- Shaozhuang Shi, Huawei
- Hailong Chu, Huawei

## [](#_description)Description

The rays to trace may be sparse in some use cases. For example, the scene only have a few regions to reflect. Providing an invocation mask image to the ray tracing commands could potentially give the hardware the hint to do certain optimization without invoking an additional pass to compact the ray buffer.

## [](#_new_commands)New Commands

- [vkCmdBindInvocationMaskHUAWEI](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBindInvocationMaskHUAWEI.html)

## [](#_new_structures)New Structures

- Extending [VkPhysicalDeviceFeatures2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceFeatures2.html), [VkDeviceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceCreateInfo.html):
  
  - [VkPhysicalDeviceInvocationMaskFeaturesHUAWEI](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceInvocationMaskFeaturesHUAWEI.html)

## [](#_new_enum_constants)New Enum Constants

- `VK_HUAWEI_INVOCATION_MASK_EXTENSION_NAME`
- `VK_HUAWEI_INVOCATION_MASK_SPEC_VERSION`
- Extending [VkAccessFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccessFlagBits2.html):
  
  - `VK_ACCESS_2_INVOCATION_MASK_READ_BIT_HUAWEI`
- Extending [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html):
  
  - `VK_IMAGE_USAGE_INVOCATION_MASK_BIT_HUAWEI`
- Extending [VkPipelineStageFlagBits2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineStageFlagBits2.html):
  
  - `VK_PIPELINE_STAGE_2_INVOCATION_MASK_BIT_HUAWEI`
- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html):
  
  - `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_INVOCATION_MASK_FEATURES_HUAWEI`

## [](#_examples)Examples

RT mask is updated before each traceRay.

Step 1. Generate InvocationMask.

```c
//the rt mask image bind as color attachment in the fragment shader
Layout(location = 2) out vec4 outRTmask
vec4 mask = vec4(x,x,x,x);
outRTmask = mask;
```

Step 2. traceRay with InvocationMask

```c
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

## [](#_version_history)Version History

- Revision 1, 2021-05-27 (Yunpeng Zhu)
  
  - Initial draft.

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VK_HUAWEI_invocation_mask).

This page is a generated document. Fixes and changes should be made to the generator scripts, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0