# VkBindPipelineIndirectCommandNV(3) Manual Page

## Name

VkBindPipelineIndirectCommandNV - Structure specifying input data for
the compute pipeline dispatch token



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkBindPipelineIndirectCommandNV` structure specifies the input data
for the `VK_INDIRECT_COMMANDS_TOKEN_TYPE_PIPELINE_NV` token.

``` c
// Provided by VK_NV_device_generated_commands_compute
typedef struct VkBindPipelineIndirectCommandNV {
    VkDeviceAddress    pipelineAddress;
} VkBindPipelineIndirectCommandNV;
```

## <a href="#_members" class="anchor"></a>Members

- `pipelineAddress` specifies the pipeline address of the compute
  pipeline that will be used in device generated rendering.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkBindPipelineIndirectCommandNV-deviceGeneratedComputePipelines-09091"
  id="VUID-VkBindPipelineIndirectCommandNV-deviceGeneratedComputePipelines-09091"></a>
  VUID-VkBindPipelineIndirectCommandNV-deviceGeneratedComputePipelines-09091  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceGeneratedComputePipelines"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV</code>::<code>deviceGeneratedComputePipelines</code></a>
  feature **must** be enabled

- <a href="#VUID-VkBindPipelineIndirectCommandNV-None-09092"
  id="VUID-VkBindPipelineIndirectCommandNV-None-09092"></a>
  VUID-VkBindPipelineIndirectCommandNV-None-09092  
  The referenced pipeline **must** have been created with
  `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`

- <a href="#VUID-VkBindPipelineIndirectCommandNV-None-09093"
  id="VUID-VkBindPipelineIndirectCommandNV-None-09093"></a>
  VUID-VkBindPipelineIndirectCommandNV-None-09093  
  The referenced pipeline **must** have been updated with
  [vkCmdUpdatePipelineIndirectBufferNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdUpdatePipelineIndirectBufferNV.html)

- <a href="#VUID-VkBindPipelineIndirectCommandNV-None-09094"
  id="VUID-VkBindPipelineIndirectCommandNV-None-09094"></a>
  VUID-VkBindPipelineIndirectCommandNV-None-09094  
  The referenced pipeline’s address **must** have been queried with
  [vkGetPipelineIndirectDeviceAddressNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPipelineIndirectDeviceAddressNV.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands_compute](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands_compute.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBindPipelineIndirectCommandNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
