# vkCreateIndirectCommandsLayoutNV(3) Manual Page

## Name

vkCreateIndirectCommandsLayoutNV - Create an indirect command layout
object



## <a href="#_c_specification" class="anchor"></a>C Specification

Indirect command layouts are created by:

``` c
// Provided by VK_NV_device_generated_commands
VkResult vkCreateIndirectCommandsLayoutNV(
    VkDevice                                    device,
    const VkIndirectCommandsLayoutCreateInfoNV* pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkIndirectCommandsLayoutNV*                 pIndirectCommandsLayout);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that creates the indirect command
  layout.

- `pCreateInfo` is a pointer to a
  [VkIndirectCommandsLayoutCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutCreateInfoNV.html)
  structure containing parameters affecting creation of the indirect
  command layout.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pIndirectCommandsLayout` is a pointer to a
  `VkIndirectCommandsLayoutNV` handle in which the resulting indirect
  command layout is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-vkCreateIndirectCommandsLayoutNV-deviceGeneratedCommands-02929"
  id="VUID-vkCreateIndirectCommandsLayoutNV-deviceGeneratedCommands-02929"></a>
  VUID-vkCreateIndirectCommandsLayoutNV-deviceGeneratedCommands-02929  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceGeneratedCommands"
  target="_blank"
  rel="noopener"><code>VkPhysicalDeviceDeviceGeneratedCommandsFeaturesNV</code>::<code>deviceGeneratedCommands</code></a>
  feature **must** be enabled

Valid Usage (Implicit)

- <a href="#VUID-vkCreateIndirectCommandsLayoutNV-device-parameter"
  id="VUID-vkCreateIndirectCommandsLayoutNV-device-parameter"></a>
  VUID-vkCreateIndirectCommandsLayoutNV-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkCreateIndirectCommandsLayoutNV-pCreateInfo-parameter"
  id="VUID-vkCreateIndirectCommandsLayoutNV-pCreateInfo-parameter"></a>
  VUID-vkCreateIndirectCommandsLayoutNV-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkIndirectCommandsLayoutCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutCreateInfoNV.html)
  structure

- <a href="#VUID-vkCreateIndirectCommandsLayoutNV-pAllocator-parameter"
  id="VUID-vkCreateIndirectCommandsLayoutNV-pAllocator-parameter"></a>
  VUID-vkCreateIndirectCommandsLayoutNV-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a
  href="#VUID-vkCreateIndirectCommandsLayoutNV-pIndirectCommandsLayout-parameter"
  id="VUID-vkCreateIndirectCommandsLayoutNV-pIndirectCommandsLayout-parameter"></a>
  VUID-vkCreateIndirectCommandsLayoutNV-pIndirectCommandsLayout-parameter  
  `pIndirectCommandsLayout` **must** be a valid pointer to a
  [VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutNV.html) handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_device_generated_commands](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_device_generated_commands.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkIndirectCommandsLayoutCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutCreateInfoNV.html),
[VkIndirectCommandsLayoutNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndirectCommandsLayoutNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateIndirectCommandsLayoutNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
