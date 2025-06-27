# vkGetGeneratedCommandsMemoryRequirementsNV(3) Manual Page

## Name

vkGetGeneratedCommandsMemoryRequirementsNV - Retrieve the buffer
allocation requirements for generated commands



## <a href="#_c_specification" class="anchor"></a>C Specification

The generation of commands on the device requires a `preprocess` buffer.
To retrieve the memory size and alignment requirements of a particular
execution state call:

``` c
// Provided by VK_NV_device_generated_commands
void vkGetGeneratedCommandsMemoryRequirementsNV(
    VkDevice                                    device,
    const VkGeneratedCommandsMemoryRequirementsInfoNV* pInfo,
    VkMemoryRequirements2*                      pMemoryRequirements);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the buffer.

- `pInfo` is a pointer to a
  [VkGeneratedCommandsMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeneratedCommandsMemoryRequirementsInfoNV.html)
  structure containing parameters required for the memory requirements
  query.

- `pMemoryRequirements` is a pointer to a
  [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html) structure in which
  the memory requirements of the buffer object are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-vkGetGeneratedCommandsMemoryRequirementsNV-deviceGeneratedCommands-02906"
  id="VUID-vkGetGeneratedCommandsMemoryRequirementsNV-deviceGeneratedCommands-02906"></a>
  VUID-vkGetGeneratedCommandsMemoryRequirementsNV-deviceGeneratedCommands-02906  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceGeneratedCommands"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceDeviceGeneratedCommandsFeaturesNV</code>::<code>deviceGeneratedCommands</code></a>
  feature **must** be enabled

- <a href="#VUID-vkGetGeneratedCommandsMemoryRequirementsNV-pInfo-09074"
  id="VUID-vkGetGeneratedCommandsMemoryRequirementsNV-pInfo-09074"></a>
  VUID-vkGetGeneratedCommandsMemoryRequirementsNV-pInfo-09074  
  If `pInfo->pipelineBindPoint` is of type
  `VK_PIPELINE_BIND_POINT_COMPUTE`, then the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceGeneratedCompute"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceDeviceGeneratedCommandsComputeFeaturesNV</code>::<code>deviceGeneratedCompute</code></a>
  feature **must** be enabled

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetGeneratedCommandsMemoryRequirementsNV-device-parameter"
  id="VUID-vkGetGeneratedCommandsMemoryRequirementsNV-device-parameter"></a>
  VUID-vkGetGeneratedCommandsMemoryRequirementsNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a
  href="#VUID-vkGetGeneratedCommandsMemoryRequirementsNV-pInfo-parameter"
  id="VUID-vkGetGeneratedCommandsMemoryRequirementsNV-pInfo-parameter"></a>
  VUID-vkGetGeneratedCommandsMemoryRequirementsNV-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid
  [VkGeneratedCommandsMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeneratedCommandsMemoryRequirementsInfoNV.html)
  structure

- <a
  href="#VUID-vkGetGeneratedCommandsMemoryRequirementsNV-pMemoryRequirements-parameter"
  id="VUID-vkGetGeneratedCommandsMemoryRequirementsNV-pMemoryRequirements-parameter"></a>
  VUID-vkGetGeneratedCommandsMemoryRequirementsNV-pMemoryRequirements-parameter  
  `pMemoryRequirements` **must** be a valid pointer to a
  [VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkGeneratedCommandsMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeneratedCommandsMemoryRequirementsInfoNV.html),
[VkMemoryRequirements2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryRequirements2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetGeneratedCommandsMemoryRequirementsNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
