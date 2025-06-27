# VkDisplaySurfaceCreateInfoKHR(3) Manual Page

## Name

VkDisplaySurfaceCreateInfoKHR - Structure specifying parameters of a
newly created display plane surface object



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDisplaySurfaceCreateInfoKHR` structure is defined as:

``` c
// Provided by VK_KHR_display
typedef struct VkDisplaySurfaceCreateInfoKHR {
    VkStructureType                   sType;
    const void*                       pNext;
    VkDisplaySurfaceCreateFlagsKHR    flags;
    VkDisplayModeKHR                  displayMode;
    uint32_t                          planeIndex;
    uint32_t                          planeStackIndex;
    VkSurfaceTransformFlagBitsKHR     transform;
    float                             globalAlpha;
    VkDisplayPlaneAlphaFlagBitsKHR    alphaMode;
    VkExtent2D                        imageExtent;
} VkDisplaySurfaceCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use, and **must** be zero.

- `displayMode` is a [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeKHR.html) handle
  specifying the mode to use when displaying this surface.

- `planeIndex` is the plane on which this surface appears.

- `planeStackIndex` is the z-order of the plane.

- `transform` is a
  [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html)
  value specifying the transformation to apply to images as part of the
  scanout operation.

- `globalAlpha` is the global alpha value. This value is ignored if
  `alphaMode` is not `VK_DISPLAY_PLANE_ALPHA_GLOBAL_BIT_KHR`.

- `alphaMode` is a
  [VkDisplayPlaneAlphaFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneAlphaFlagBitsKHR.html)
  value specifying the type of alpha blending to use.

- `imageExtent` is the size of the presentable images to use with the
  surface.

## <a href="#_description" class="anchor"></a>Description

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Creating a display surface <strong>must</strong> not modify the state
of the displays, planes, or other resources it names. For example, it
<strong>must</strong> not apply the specified mode to be set on the
associated display. Application of display configuration occurs as a
side effect of presenting to a display surface.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-VkDisplaySurfaceCreateInfoKHR-planeIndex-01252"
  id="VUID-VkDisplaySurfaceCreateInfoKHR-planeIndex-01252"></a>
  VUID-VkDisplaySurfaceCreateInfoKHR-planeIndex-01252  
  `planeIndex` **must** be less than the number of display planes
  supported by the device as determined by calling
  `vkGetPhysicalDeviceDisplayPlanePropertiesKHR`

- <a href="#VUID-VkDisplaySurfaceCreateInfoKHR-planeReorderPossible-01253"
  id="VUID-VkDisplaySurfaceCreateInfoKHR-planeReorderPossible-01253"></a>
  VUID-VkDisplaySurfaceCreateInfoKHR-planeReorderPossible-01253  
  If the `planeReorderPossible` member of the `VkDisplayPropertiesKHR`
  structure returned by `vkGetPhysicalDeviceDisplayPropertiesKHR` for
  the display corresponding to `displayMode` is `VK_TRUE` then
  `planeStackIndex` **must** be less than the number of display planes
  supported by the device as determined by calling
  `vkGetPhysicalDeviceDisplayPlanePropertiesKHR`; otherwise
  `planeStackIndex` **must** equal the `currentStackIndex` member of
  `VkDisplayPlanePropertiesKHR` returned by
  `vkGetPhysicalDeviceDisplayPlanePropertiesKHR` for the display plane
  corresponding to `displayMode`

- <a href="#VUID-VkDisplaySurfaceCreateInfoKHR-alphaMode-01254"
  id="VUID-VkDisplaySurfaceCreateInfoKHR-alphaMode-01254"></a>
  VUID-VkDisplaySurfaceCreateInfoKHR-alphaMode-01254  
  If `alphaMode` is `VK_DISPLAY_PLANE_ALPHA_GLOBAL_BIT_KHR` then
  `globalAlpha` **must** be between `0` and `1`, inclusive

- <a href="#VUID-VkDisplaySurfaceCreateInfoKHR-alphaMode-01255"
  id="VUID-VkDisplaySurfaceCreateInfoKHR-alphaMode-01255"></a>
  VUID-VkDisplaySurfaceCreateInfoKHR-alphaMode-01255  
  `alphaMode` **must** be one of the bits present in the
  `supportedAlpha` member of `VkDisplayPlaneCapabilitiesKHR` for the
  display plane corresponding to `displayMode`

- <a href="#VUID-VkDisplaySurfaceCreateInfoKHR-transform-06740"
  id="VUID-VkDisplaySurfaceCreateInfoKHR-transform-06740"></a>
  VUID-VkDisplaySurfaceCreateInfoKHR-transform-06740  
  `transform` **must** be one of the bits present in the
  `supportedTransforms` member of `VkDisplayPropertiesKHR` for the
  display corresponding to `displayMode`

- <a href="#VUID-VkDisplaySurfaceCreateInfoKHR-width-01256"
  id="VUID-VkDisplaySurfaceCreateInfoKHR-width-01256"></a>
  VUID-VkDisplaySurfaceCreateInfoKHR-width-01256  
  The `width` and `height` members of `imageExtent` **must** be less
  than or equal to
  [VkPhysicalDeviceLimits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLimits.html)::`maxImageDimension2D`

Valid Usage (Implicit)

- <a href="#VUID-VkDisplaySurfaceCreateInfoKHR-sType-sType"
  id="VUID-VkDisplaySurfaceCreateInfoKHR-sType-sType"></a>
  VUID-VkDisplaySurfaceCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DISPLAY_SURFACE_CREATE_INFO_KHR`

- <a href="#VUID-VkDisplaySurfaceCreateInfoKHR-pNext-pNext"
  id="VUID-VkDisplaySurfaceCreateInfoKHR-pNext-pNext"></a>
  VUID-VkDisplaySurfaceCreateInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkDisplaySurfaceCreateInfoKHR-flags-zerobitmask"
  id="VUID-VkDisplaySurfaceCreateInfoKHR-flags-zerobitmask"></a>
  VUID-VkDisplaySurfaceCreateInfoKHR-flags-zerobitmask  
  `flags` **must** be `0`

- <a href="#VUID-VkDisplaySurfaceCreateInfoKHR-displayMode-parameter"
  id="VUID-VkDisplaySurfaceCreateInfoKHR-displayMode-parameter"></a>
  VUID-VkDisplaySurfaceCreateInfoKHR-displayMode-parameter  
  `displayMode` **must** be a valid
  [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeKHR.html) handle

- <a href="#VUID-VkDisplaySurfaceCreateInfoKHR-transform-parameter"
  id="VUID-VkDisplaySurfaceCreateInfoKHR-transform-parameter"></a>
  VUID-VkDisplaySurfaceCreateInfoKHR-transform-parameter  
  `transform` **must** be a valid
  [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html)
  value

- <a href="#VUID-VkDisplaySurfaceCreateInfoKHR-alphaMode-parameter"
  id="VUID-VkDisplaySurfaceCreateInfoKHR-alphaMode-parameter"></a>
  VUID-VkDisplaySurfaceCreateInfoKHR-alphaMode-parameter  
  `alphaMode` **must** be a valid
  [VkDisplayPlaneAlphaFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneAlphaFlagBitsKHR.html)
  value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html),
[VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeKHR.html),
[VkDisplayPlaneAlphaFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneAlphaFlagBitsKHR.html),
[VkDisplaySurfaceCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplaySurfaceCreateFlagsKHR.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html),
[vkCreateDisplayPlaneSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateDisplayPlaneSurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDisplaySurfaceCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
