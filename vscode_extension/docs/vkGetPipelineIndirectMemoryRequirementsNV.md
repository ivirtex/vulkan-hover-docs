# vkGetPipelineIndirectMemoryRequirementsNV(3) Manual Page

## Name

vkGetPipelineIndirectMemoryRequirementsNV - Get the memory requirements
for the compute indirect pipeline



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine the memory requirements for a compute pipeline’s metadata,
call:

``` c
// Provided by VK_NV_device_generated_commands_compute
void vkGetPipelineIndirectMemoryRequirementsNV(
    VkDevice                                    device,
    const VkComputePipelineCreateInfo*          pCreateInfo,
    VkMemoryRequirements2*                      pMemoryRequirements);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the buffer.

- `pCreateInfo` is a
  [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html)
  structure specifying the creation parameters of the compute pipeline
  whose memory requirements are being queried.

- `pMemoryRequirements` is a pointer to a
  [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html) structure in which
  the requested pipeline’s memory requirements are returned.

## <a href="#_description" class="anchor"></a>Description

If `pCreateInfo->pNext` chain includes a pointer to a
[VkComputePipelineIndirectBufferInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineIndirectBufferInfoNV.html)
structure, then the contents of that structure are ignored.

Valid Usage

- <a
  href="#VUID-vkGetPipelineIndirectMemoryRequirementsNV-deviceGeneratedComputePipelines-09082"
  id="VUID-vkGetPipelineIndirectMemoryRequirementsNV-deviceGeneratedComputePipelines-09082"></a>
  VUID-vkGetPipelineIndirectMemoryRequirementsNV-deviceGeneratedComputePipelines-09082  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceGeneratedComputePipelines"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV</code>::<code>deviceGeneratedComputePipelines</code></a>
  feature **must** be enabled

- <a
  href="#VUID-vkGetPipelineIndirectMemoryRequirementsNV-pCreateInfo-09083"
  id="VUID-vkGetPipelineIndirectMemoryRequirementsNV-pCreateInfo-09083"></a>
  VUID-vkGetPipelineIndirectMemoryRequirementsNV-pCreateInfo-09083  
  `pCreateInfo->flags` **must** include
  `VK_PIPELINE_CREATE_INDIRECT_BINDABLE_BIT_NV`

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPipelineIndirectMemoryRequirementsNV-device-parameter"
  id="VUID-vkGetPipelineIndirectMemoryRequirementsNV-device-parameter"></a>
  VUID-vkGetPipelineIndirectMemoryRequirementsNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetPipelineIndirectMemoryRequirementsNV-pCreateInfo-parameter"
  id="VUID-vkGetPipelineIndirectMemoryRequirementsNV-pCreateInfo-parameter"></a>
  VUID-vkGetPipelineIndirectMemoryRequirementsNV-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html)
  structure

- <a
  href="#VUID-vkGetPipelineIndirectMemoryRequirementsNV-pMemoryRequirements-parameter"
  id="VUID-vkGetPipelineIndirectMemoryRequirementsNV-pMemoryRequirements-parameter"></a>
  VUID-vkGetPipelineIndirectMemoryRequirementsNV-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a
  [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands_compute](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands_compute.html),
[VkComputePipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComputePipelineCreateInfo.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPipelineIndirectMemoryRequirementsNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
