# vkGetPhysicalDevicePresentRectanglesKHR(3) Manual Page

## Name

vkGetPhysicalDevicePresentRectanglesKHR - Query present rectangles for a
surface on a physical device



## <a href="#_c_specification" class="anchor"></a>C Specification

When using `VK_DEVICE_GROUP_PRESENT_MODE_LOCAL_MULTI_DEVICE_BIT_KHR`,
the application **may** need to know which regions of the surface are
used when presenting locally on each physical device. Presentation of
swapchain images to this surface need only have valid contents in the
regions returned by this command.

To query a set of rectangles used in presentation on the physical
device, call:

``` c
// Provided by VK_VERSION_1_1 with VK_KHR_swapchain, VK_KHR_device_group with VK_KHR_surface
VkResult vkGetPhysicalDevicePresentRectanglesKHR(
    VkPhysicalDevice                            physicalDevice,
    VkSurfaceKHR                                surface,
    uint32_t*                                   pRectCount,
    VkRect2D*                                   pRects);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device.

- `surface` is the surface.

- `pRectCount` is a pointer to an integer related to the number of
  rectangles available or queried, as described below.

- `pRects` is either `NULL` or a pointer to an array of
  [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) structures.

## <a href="#_description" class="anchor"></a>Description

If `pRects` is `NULL`, then the number of rectangles used when
presenting the given `surface` is returned in `pRectCount`. Otherwise,
`pRectCount` **must** point to a variable set by the application to the
number of elements in the `pRects` array, and on return the variable is
overwritten with the number of structures actually written to `pRects`.
If the value of `pRectCount` is less than the number of rectangles, at
most `pRectCount` structures will be written, and `VK_INCOMPLETE` will
be returned instead of `VK_SUCCESS`, to indicate that not all the
available rectangles were returned.

The values returned by this command are not invariant, and **may**
change in response to the surface being moved, resized, or occluded.

The rectangles returned by this command **must** not overlap.

Valid Usage

- <a href="#VUID-vkGetPhysicalDevicePresentRectanglesKHR-surface-06523"
  id="VUID-vkGetPhysicalDevicePresentRectanglesKHR-surface-06523"></a>
  VUID-vkGetPhysicalDevicePresentRectanglesKHR-surface-06523  
  `surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

- <a href="#VUID-vkGetPhysicalDevicePresentRectanglesKHR-surface-06211"
  id="VUID-vkGetPhysicalDevicePresentRectanglesKHR-surface-06211"></a>
  VUID-vkGetPhysicalDevicePresentRectanglesKHR-surface-06211  
  `surface` **must** be supported by `physicalDevice`, as reported by
  [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html)
  or an equivalent platform-specific mechanism

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDevicePresentRectanglesKHR-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDevicePresentRectanglesKHR-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDevicePresentRectanglesKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDevicePresentRectanglesKHR-surface-parameter"
  id="VUID-vkGetPhysicalDevicePresentRectanglesKHR-surface-parameter"></a>
  VUID-vkGetPhysicalDevicePresentRectanglesKHR-surface-parameter  
  `surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

- <a
  href="#VUID-vkGetPhysicalDevicePresentRectanglesKHR-pRectCount-parameter"
  id="VUID-vkGetPhysicalDevicePresentRectanglesKHR-pRectCount-parameter"></a>
  VUID-vkGetPhysicalDevicePresentRectanglesKHR-pRectCount-parameter  
  `pRectCount` **must** be a valid pointer to a `uint32_t` value

- <a href="#VUID-vkGetPhysicalDevicePresentRectanglesKHR-pRects-parameter"
  id="VUID-vkGetPhysicalDevicePresentRectanglesKHR-pRects-parameter"></a>
  VUID-vkGetPhysicalDevicePresentRectanglesKHR-pRects-parameter  
  If the value referenced by `pRectCount` is not `0`, and `pRects` is
  not `NULL`, `pRects` **must** be a valid pointer to an array of
  `pRectCount` [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html) structures

- <a href="#VUID-vkGetPhysicalDevicePresentRectanglesKHR-commonparent"
  id="VUID-vkGetPhysicalDevicePresentRectanglesKHR-commonparent"></a>
  VUID-vkGetPhysicalDevicePresentRectanglesKHR-commonparent  
  Both of `physicalDevice`, and `surface` **must** have been created,
  allocated, or retrieved from the same [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html)

Host Synchronization

- Host access to `surface` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_device_group](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group.html),
[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html),
[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html),
[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html), [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html),
[VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDevicePresentRectanglesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
