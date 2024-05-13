# vkGetPipelineIndirectDeviceAddressNV(3) Manual Page

## Name

vkGetPipelineIndirectDeviceAddressNV - Get pipeline's 64-bit device
address



## <a href="#_c_specification" class="anchor"></a>C Specification

To query a compute pipeline’s 64-bit device address, call:

``` c
// Provided by VK_NV_device_generated_commands_compute
VkDeviceAddress vkGetPipelineIndirectDeviceAddressNV(
    VkDevice                                    device,
    const VkPipelineIndirectDeviceAddressInfoNV* pInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device on which the pipeline was created.

- `pInfo` is a pointer to a
  [VkPipelineIndirectDeviceAddressInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineIndirectDeviceAddressInfoNV.html)
  structure specifying the pipeline to retrieve the address for.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-vkGetPipelineIndirectDeviceAddressNV-deviceGeneratedComputePipelines-09078"
  id="VUID-vkGetPipelineIndirectDeviceAddressNV-deviceGeneratedComputePipelines-09078"></a>
  VUID-vkGetPipelineIndirectDeviceAddressNV-deviceGeneratedComputePipelines-09078  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceGeneratedComputePipelines"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV</code>::<code>deviceGeneratedComputePipelines</code></a>
  feature **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkGetPipelineIndirectDeviceAddressNV-device-parameter"
  id="VUID-vkGetPipelineIndirectDeviceAddressNV-device-parameter"></a>
  VUID-vkGetPipelineIndirectDeviceAddressNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetPipelineIndirectDeviceAddressNV-pInfo-parameter"
  id="VUID-vkGetPipelineIndirectDeviceAddressNV-pInfo-parameter"></a>
  VUID-vkGetPipelineIndirectDeviceAddressNV-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkPipelineIndirectDeviceAddressInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineIndirectDeviceAddressInfoNV.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands_compute](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands_compute.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkPipelineIndirectDeviceAddressInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineIndirectDeviceAddressInfoNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPipelineIndirectDeviceAddressNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
