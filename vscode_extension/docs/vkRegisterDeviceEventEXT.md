# vkRegisterDeviceEventEXT(3) Manual Page

## Name

vkRegisterDeviceEventEXT - Signal a fence when a device event occurs



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a fence that will be signaled when an event occurs on a
device, call:

``` c
// Provided by VK_EXT_display_control
VkResult vkRegisterDeviceEventEXT(
    VkDevice                                    device,
    const VkDeviceEventInfoEXT*                 pDeviceEventInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkFence*                                    pFence);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is a logical device on which the event **may** occur.

- `pDeviceEventInfo` is a pointer to a
  [VkDeviceEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceEventInfoEXT.html) structure describing
  the event of interest to the application.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pFence` is a pointer to a handle in which the resulting fence object
  is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkRegisterDeviceEventEXT-device-parameter"
  id="VUID-vkRegisterDeviceEventEXT-device-parameter"></a>
  VUID-vkRegisterDeviceEventEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkRegisterDeviceEventEXT-pDeviceEventInfo-parameter"
  id="VUID-vkRegisterDeviceEventEXT-pDeviceEventInfo-parameter"></a>
  VUID-vkRegisterDeviceEventEXT-pDeviceEventInfo-parameter  
  `pDeviceEventInfo` **must** be a valid pointer to a valid
  [VkDeviceEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceEventInfoEXT.html) structure

- <a href="#VUID-vkRegisterDeviceEventEXT-pAllocator-parameter"
  id="VUID-vkRegisterDeviceEventEXT-pAllocator-parameter"></a>
  VUID-vkRegisterDeviceEventEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkRegisterDeviceEventEXT-pFence-parameter"
  id="VUID-vkRegisterDeviceEventEXT-pFence-parameter"></a>
  VUID-vkRegisterDeviceEventEXT-pFence-parameter  
  `pFence` **must** be a valid pointer to a [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html)
  handle

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_display_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_display_control.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkDeviceEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceEventInfoEXT.html),
[VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkRegisterDeviceEventEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
