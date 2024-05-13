# vkCreateDisplayModeKHR(3) Manual Page

## Name

vkCreateDisplayModeKHR - Create a display mode



## <a href="#_c_specification" class="anchor"></a>C Specification

Additional modes **may** also be created by calling:

``` c
// Provided by VK_KHR_display
VkResult vkCreateDisplayModeKHR(
    VkPhysicalDevice                            physicalDevice,
    VkDisplayKHR                                display,
    const VkDisplayModeCreateInfoKHR*           pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDisplayModeKHR*                           pMode);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device associated with `display`.

- `display` is the display to create an additional mode for.

- `pCreateInfo` is a pointer to a
  [VkDisplayModeCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeCreateInfoKHR.html)
  structure describing the new mode to create.

- `pAllocator` is the allocator used for host memory allocated for the
  display mode object when there is no more specific allocator available
  (see <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a>).

- `pMode` is a pointer to a [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeKHR.html)
  handle in which the mode created is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkCreateDisplayModeKHR-physicalDevice-parameter"
  id="VUID-vkCreateDisplayModeKHR-physicalDevice-parameter"></a>
  VUID-vkCreateDisplayModeKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkCreateDisplayModeKHR-display-parameter"
  id="VUID-vkCreateDisplayModeKHR-display-parameter"></a>
  VUID-vkCreateDisplayModeKHR-display-parameter  
  `display` **must** be a valid [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) handle

- <a href="#VUID-vkCreateDisplayModeKHR-pCreateInfo-parameter"
  id="VUID-vkCreateDisplayModeKHR-pCreateInfo-parameter"></a>
  VUID-vkCreateDisplayModeKHR-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid
  [VkDisplayModeCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeCreateInfoKHR.html)
  structure

- <a href="#VUID-vkCreateDisplayModeKHR-pAllocator-parameter"
  id="VUID-vkCreateDisplayModeKHR-pAllocator-parameter"></a>
  VUID-vkCreateDisplayModeKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkCreateDisplayModeKHR-pMode-parameter"
  id="VUID-vkCreateDisplayModeKHR-pMode-parameter"></a>
  VUID-vkCreateDisplayModeKHR-pMode-parameter  
  `pMode` **must** be a valid pointer to a
  [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeKHR.html) handle

- <a href="#VUID-vkCreateDisplayModeKHR-display-parent"
  id="VUID-vkCreateDisplayModeKHR-display-parent"></a>
  VUID-vkCreateDisplayModeKHR-display-parent  
  `display` **must** have been created, allocated, or retrieved from
  `physicalDevice`

Host Synchronization

- Host access to `display` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_INITIALIZATION_FAILED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html),
[VkDisplayModeCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeCreateInfoKHR.html),
[VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeKHR.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCreateDisplayModeKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
