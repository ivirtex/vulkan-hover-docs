# vkGetPipelineIndirectDeviceAddressNV(3) Manual Page

## Name

vkGetPipelineIndirectDeviceAddressNV - Get pipeline's 64-bit device address



## [](#_c_specification)C Specification

To query a compute pipeline’s 64-bit device address, call:

```c++
// Provided by VK_NV_device_generated_commands_compute
VkDeviceAddress vkGetPipelineIndirectDeviceAddressNV(
    VkDevice                                    device,
    const VkPipelineIndirectDeviceAddressInfoNV* pInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device on which the pipeline was created.
- `pInfo` is a pointer to a [VkPipelineIndirectDeviceAddressInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineIndirectDeviceAddressInfoNV.html) structure specifying the pipeline to retrieve the address for.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetPipelineIndirectDeviceAddressNV-deviceGeneratedComputePipelines-09078)VUID-vkGetPipelineIndirectDeviceAddressNV-deviceGeneratedComputePipelines-09078  
  The [`VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV`::`deviceGeneratedComputePipelines`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedComputePipelines) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetPipelineIndirectDeviceAddressNV-device-parameter)VUID-vkGetPipelineIndirectDeviceAddressNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetPipelineIndirectDeviceAddressNV-pInfo-parameter)VUID-vkGetPipelineIndirectDeviceAddressNV-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkPipelineIndirectDeviceAddressInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineIndirectDeviceAddressInfoNV.html) structure

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands\_compute](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands_compute.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkPipelineIndirectDeviceAddressInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipelineIndirectDeviceAddressInfoNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPipelineIndirectDeviceAddressNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0