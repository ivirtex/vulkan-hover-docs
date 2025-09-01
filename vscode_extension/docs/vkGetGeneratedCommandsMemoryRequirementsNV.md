# vkGetGeneratedCommandsMemoryRequirementsNV(3) Manual Page

## Name

vkGetGeneratedCommandsMemoryRequirementsNV - Retrieve the buffer allocation requirements for generated commands



## [](#_c_specification)C Specification

With `VK_NV_device_generated_commands`, to retrieve the memory size and alignment requirements of a particular execution state call:

```c++
// Provided by VK_NV_device_generated_commands
void vkGetGeneratedCommandsMemoryRequirementsNV(
    VkDevice                                    device,
    const VkGeneratedCommandsMemoryRequirementsInfoNV* pInfo,
    VkMemoryRequirements2*                      pMemoryRequirements);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the buffer.
- `pInfo` is a pointer to a [VkGeneratedCommandsMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsMemoryRequirementsInfoNV.html) structure containing parameters required for the memory requirements query.
- `pMemoryRequirements` is a pointer to a [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html) structure in which the memory requirements of the buffer object are returned.

## [](#_description)Description

Valid Usage

- [](#VUID-vkGetGeneratedCommandsMemoryRequirementsNV-deviceGeneratedCommands-02906)VUID-vkGetGeneratedCommandsMemoryRequirementsNV-deviceGeneratedCommands-02906  
  The [`VkPhysicalDeviceDeviceGeneratedCommandsFeaturesNV`::`deviceGeneratedCommands`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedCommandsNV) feature **must** be enabled
- [](#VUID-vkGetGeneratedCommandsMemoryRequirementsNV-pInfo-09074)VUID-vkGetGeneratedCommandsMemoryRequirementsNV-pInfo-09074  
  If `pInfo->pipelineBindPoint` is of type `VK_PIPELINE_BIND_POINT_COMPUTE`, then the [`VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV`::`deviceGeneratedCompute`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceGeneratedCompute) feature **must** be enabled

Valid Usage (Implicit)

- [](#VUID-vkGetGeneratedCommandsMemoryRequirementsNV-device-parameter)VUID-vkGetGeneratedCommandsMemoryRequirementsNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetGeneratedCommandsMemoryRequirementsNV-pInfo-parameter)VUID-vkGetGeneratedCommandsMemoryRequirementsNV-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkGeneratedCommandsMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsMemoryRequirementsInfoNV.html) structure
- [](#VUID-vkGetGeneratedCommandsMemoryRequirementsNV-pMemoryRequirements-parameter)VUID-vkGetGeneratedCommandsMemoryRequirementsNV-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html) structure

## [](#_see_also)See Also

[VK\_NV\_device\_generated\_commands](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_device_generated_commands.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkGeneratedCommandsMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkGeneratedCommandsMemoryRequirementsInfoNV.html), [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetGeneratedCommandsMemoryRequirementsNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0