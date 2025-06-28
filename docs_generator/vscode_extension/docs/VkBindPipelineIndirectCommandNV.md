# VkBindPipelineIndirectCommandNV(3) Manual Page

## Name

VkBindPipelineIndirectCommandNV - Structure specifying input data for the compute pipeline dispatch token



## [](#_c_specification)C Specification

The `VkBindPipelineIndirectCommandNV` structure specifies the input data for the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV` token.

```c++
// Provided by VK_NV_device_generated_commands_compute
typedef struct VkBindPipelineIndirectCommandNV {
    VkDeviceAddress    pipelineAddress;
} VkBindPipelineIndirectCommandNV;
```

## [](#_members)Members

- `pipelineAddress` specifies the pipeline address of the compute pipeline that will be used in device generated rendering.

## [](#_description)Description

Valid Usage

- [](#VUID-VkBindPipelineIndirectCommandNV-deviceGeneratedComputePipelines-09091)VUID-VkBindPipelineIndirectCommandNV-deviceGeneratedComputePipelines-09091  
  The [`VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV`::`deviceGeneratedComputePipelines`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedComputePipelines) feature **must** be enabled
- [](#VUID-VkBindPipelineIndirectCommandNV-None-09092)VUID-VkBindPipelineIndirectCommandNV-None-09092  
  The referenced pipeline **must** have been created with `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`
- [](#VUID-VkBindPipelineIndirectCommandNV-None-09093)VUID-VkBindPipelineIndirectCommandNV-None-09093  
  The referenced pipeline **must** have been updated with [vkCmdUpdatePipelineIndirectBufferNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdUpdatePipelineIndirectBufferNV.html)
- [](#VUID-VkBindPipelineIndirectCommandNV-None-09094)VUID-VkBindPipelineIndirectCommandNV-None-09094  
  The referenced pipeline’s address **must** have been queried with [vkGetPipelineIndirectDeviceAddressNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPipelineIndirectDeviceAddressNV.html)

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands\_compute](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands_compute.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBindPipelineIndirectCommandNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0