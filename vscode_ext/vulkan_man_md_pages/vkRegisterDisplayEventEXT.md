# vkRegisterDisplayEventEXT(3) Manual Page

## Name

vkRegisterDisplayEventEXT - Signal a fence when a display event occurs



## <a href="#_c_specification" class="anchor"></a>C Specification

To create a fence that will be signaled when an event occurs on a
[VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) object, call:

``` c
// Provided by VK_EXT_display_control
VkResult vkRegisterDisplayEventEXT(
    VkDevice                                    device,
    VkDisplayKHR                                display,
    const VkDisplayEventInfoEXT*                pDisplayEventInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkFence*                                    pFence);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is a logical device associated with `display`

- `display` is the display on which the event **may** occur.

- `pDisplayEventInfo` is a pointer to a
  [VkDisplayEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayEventInfoEXT.html) structure
  describing the event of interest to the application.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

- `pFence` is a pointer to a handle in which the resulting fence object
  is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkRegisterDisplayEventEXT-device-parameter"
  id="VUID-vkRegisterDisplayEventEXT-device-parameter"></a>
  VUID-vkRegisterDisplayEventEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkRegisterDisplayEventEXT-display-parameter"
  id="VUID-vkRegisterDisplayEventEXT-display-parameter"></a>
  VUID-vkRegisterDisplayEventEXT-display-parameter  
  `display` **must** be a valid [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) handle

- <a href="#VUID-vkRegisterDisplayEventEXT-pDisplayEventInfo-parameter"
  id="VUID-vkRegisterDisplayEventEXT-pDisplayEventInfo-parameter"></a>
  VUID-vkRegisterDisplayEventEXT-pDisplayEventInfo-parameter  
  `pDisplayEventInfo` **must** be a valid pointer to a valid
  [VkDisplayEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayEventInfoEXT.html) structure

- <a href="#VUID-vkRegisterDisplayEventEXT-pAllocator-parameter"
  id="VUID-vkRegisterDisplayEventEXT-pAllocator-parameter"></a>
  VUID-vkRegisterDisplayEventEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkRegisterDisplayEventEXT-pFence-parameter"
  id="VUID-vkRegisterDisplayEventEXT-pFence-parameter"></a>
  VUID-vkRegisterDisplayEventEXT-pFence-parameter  
  `pFence` **must** be a valid pointer to a [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html)
  handle

- <a href="#VUID-vkRegisterDisplayEventEXT-commonparent"
  id="VUID-vkRegisterDisplayEventEXT-commonparent"></a>
  VUID-vkRegisterDisplayEventEXT-commonparent  
  Both of `device`, and `display` **must** have been created, allocated,
  or retrieved from the same [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_display_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_display_control.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkDisplayEventInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayEventInfoEXT.html),
[VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html), [VkFence](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFence.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkRegisterDisplayEventEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
