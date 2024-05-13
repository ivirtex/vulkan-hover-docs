# VkSwapchainPresentScalingCreateInfoEXT(3) Manual Page

## Name

VkSwapchainPresentScalingCreateInfoEXT - Scaling behavior when
presenting to the surface



## <a href="#_c_specification" class="anchor"></a>C Specification

When an application presents a swapchain image with dimensions different
than those of the target surface, different behavior is possible on
different platforms per their respective specifications:

- Presentation fails and `VK_ERROR_OUT_OF_DATE_KHR` is returned

- Scaling is done and `VK_SUCCESS` or `VK_SUBOPTIMAL_KHR` is returned

- Unspecified scaling using an arbitrary combination of stretching,
  centering and/or clipping.

Applications **can** define specific behavior when creating a swapchain
by including the `VkSwapchainPresentScalingCreateInfoEXT` structure in
the `pNext` chain of the
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html) structure.

The `VkSwapchainPresentScalingCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_swapchain_maintenance1
typedef struct VkSwapchainPresentScalingCreateInfoEXT {
    VkStructureType             sType;
    const void*                 pNext;
    VkPresentScalingFlagsEXT    scalingBehavior;
    VkPresentGravityFlagsEXT    presentGravityX;
    VkPresentGravityFlagsEXT    presentGravityY;
} VkSwapchainPresentScalingCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `scalingBehavior` is `0` or the scaling method to use when the
  dimensions of the surface and swapchain images differ.

- `presentGravityX` is `0` or the x-axis direction in which swapchain
  image pixels gravitate relative to the surface when `scalingBehavior`
  does not result in a one-to-one pixel mapping between the scaled
  swapchain image and the surface.

- `presentGravityY` is `0` or the y-axis direction in which swapchain
  image pixels gravitate relative to the surface when `scalingBehavior`
  does not result in a one-to-one pixel mapping between the scaled
  swapchain image and the surface.

## <a href="#_description" class="anchor"></a>Description

If `scalingBehavior` is `0`, the result of presenting a swapchain image
with dimensions that do not match the surface dimensions is
implementation and platform-dependent. If `presentGravityX` or
`presentGravityY` are `0`, the presentation gravity **must** match that
defined by the native platform surface on platforms which define surface
gravity.

Valid Usage

- <a
  href="#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07765"
  id="VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07765"></a>
  VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07765  
  If `presentGravityX` is `0`, `presentGravityY` **must** be `0`

- <a
  href="#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07766"
  id="VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07766"></a>
  VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07766  
  If `presentGravityX` is not `0`, `presentGravityY` **must** not be `0`

- <a
  href="#VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-07767"
  id="VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-07767"></a>
  VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-07767  
  `scalingBehavior` **must** not have more than one bit set

- <a
  href="#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07768"
  id="VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07768"></a>
  VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07768  
  `presentGravityX` **must** not have more than one bit set

- <a
  href="#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-07769"
  id="VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-07769"></a>
  VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-07769  
  `presentGravityY` **must** not have more than one bit set

- <a
  href="#VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-07770"
  id="VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-07770"></a>
  VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-07770  
  `scalingBehavior` **must** be a valid scaling method for the surface
  as returned in
  [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`supportedPresentScaling`,
  given
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html)::`presentMode`
  in [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeEXT.html)

- <a
  href="#VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-07771"
  id="VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-07771"></a>
  VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-07771  
  If the swapchain is created with
  [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentModesCreateInfoEXT.html),
  `scalingBehavior` **must** be a valid scaling method for the surface
  as returned in
  [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`supportedPresentScaling`,
  given each present mode in
  [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentModesCreateInfoEXT.html)::`pPresentModes`
  in [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeEXT.html)

- <a
  href="#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07772"
  id="VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07772"></a>
  VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07772  
  `presentGravityX` **must** be a valid x-axis present gravity for the
  surface as returned in
  [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`supportedPresentGravityX`,
  given
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html)::`presentMode`
  in [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeEXT.html)

- <a
  href="#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07773"
  id="VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07773"></a>
  VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-07773  
  If the swapchain is created with
  [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentModesCreateInfoEXT.html),
  `presentGravityX` **must** be a valid x-axis present gravity for the
  surface as returned in
  [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`supportedPresentGravityX`,
  given each present mode in
  [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentModesCreateInfoEXT.html)::`pPresentModes`
  in [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeEXT.html)

- <a
  href="#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-07774"
  id="VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-07774"></a>
  VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-07774  
  `presentGravityY` **must** be a valid y-axis present gravity for the
  surface as returned in
  [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`supportedPresentGravityY`,
  given
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html)::`presentMode`
  in [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeEXT.html)

- <a
  href="#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-07775"
  id="VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-07775"></a>
  VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-07775  
  If the swapchain is created with
  [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentModesCreateInfoEXT.html),
  `presentGravityY` **must** be a valid y-axis present gravity for the
  surface as returned in
  [VkSurfacePresentScalingCapabilitiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentScalingCapabilitiesEXT.html)::`supportedPresentGravityY`,
  given each present mode in
  [VkSwapchainPresentModesCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainPresentModesCreateInfoEXT.html)::`pPresentModes`
  in [VkSurfacePresentModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfacePresentModeEXT.html)

Valid Usage (Implicit)

- <a href="#VUID-VkSwapchainPresentScalingCreateInfoEXT-sType-sType"
  id="VUID-VkSwapchainPresentScalingCreateInfoEXT-sType-sType"></a>
  VUID-VkSwapchainPresentScalingCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SWAPCHAIN_PRESENT_SCALING_CREATE_INFO_EXT`

- <a
  href="#VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-parameter"
  id="VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-parameter"></a>
  VUID-VkSwapchainPresentScalingCreateInfoEXT-scalingBehavior-parameter  
  `scalingBehavior` **must** be a valid combination of
  [VkPresentScalingFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentScalingFlagBitsEXT.html) values

- <a
  href="#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-parameter"
  id="VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-parameter"></a>
  VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityX-parameter  
  `presentGravityX` **must** be a valid combination of
  [VkPresentGravityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentGravityFlagBitsEXT.html) values

- <a
  href="#VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-parameter"
  id="VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-parameter"></a>
  VUID-VkSwapchainPresentScalingCreateInfoEXT-presentGravityY-parameter  
  `presentGravityY` **must** be a valid combination of
  [VkPresentGravityFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentGravityFlagBitsEXT.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_swapchain_maintenance1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_swapchain_maintenance1.html),
[VkPresentGravityFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentGravityFlagsEXT.html),
[VkPresentScalingFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentScalingFlagsEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSwapchainPresentScalingCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
