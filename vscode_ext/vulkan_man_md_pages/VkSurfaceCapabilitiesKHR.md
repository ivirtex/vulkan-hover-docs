# VkSurfaceCapabilitiesKHR(3) Manual Page

## Name

VkSurfaceCapabilitiesKHR - Structure describing capabilities of a
surface



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSurfaceCapabilitiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_surface
typedef struct VkSurfaceCapabilitiesKHR {
    uint32_t                         minImageCount;
    uint32_t                         maxImageCount;
    VkExtent2D                       currentExtent;
    VkExtent2D                       minImageExtent;
    VkExtent2D                       maxImageExtent;
    uint32_t                         maxImageArrayLayers;
    VkSurfaceTransformFlagsKHR       supportedTransforms;
    VkSurfaceTransformFlagBitsKHR    currentTransform;
    VkCompositeAlphaFlagsKHR         supportedCompositeAlpha;
    VkImageUsageFlags                supportedUsageFlags;
} VkSurfaceCapabilitiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

## <a href="#_description" class="anchor"></a>Description

- `minImageCount` is the minimum number of images the specified device
  supports for a swapchain created for the surface, and will be at least
  one.

- `maxImageCount` is the maximum number of images the specified device
  supports for a swapchain created for the surface, and will be either
  0, or greater than or equal to `minImageCount`. A value of 0 means
  that there is no limit on the number of images, though there **may**
  be limits related to the total amount of memory used by presentable
  images.

- `currentExtent` is the current width and height of the surface, or the
  special value (0xFFFFFFFF, 0xFFFFFFFF) indicating that the surface
  size will be determined by the extent of a swapchain targeting the
  surface.

- `minImageExtent` contains the smallest valid swapchain extent for the
  surface on the specified device. The `width` and `height` of the
  extent will each be less than or equal to the corresponding `width`
  and `height` of `currentExtent`, unless `currentExtent` has the
  special value described above.

- `maxImageExtent` contains the largest valid swapchain extent for the
  surface on the specified device. The `width` and `height` of the
  extent will each be greater than or equal to the corresponding `width`
  and `height` of `minImageExtent`. The `width` and `height` of the
  extent will each be greater than or equal to the corresponding `width`
  and `height` of `currentExtent`, unless `currentExtent` has the
  special value described above.

- `maxImageArrayLayers` is the maximum number of layers presentable
  images **can** have for a swapchain created for this device and
  surface, and will be at least one.

- `supportedTransforms` is a bitmask of
  [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html)
  indicating the presentation transforms supported for the surface on
  the specified device. At least one bit will be set.

- `currentTransform` is
  [VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html)
  value indicating the surface’s current transform relative to the
  presentation engine’s natural orientation.

- `supportedCompositeAlpha` is a bitmask of
  [VkCompositeAlphaFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompositeAlphaFlagBitsKHR.html),
  representing the alpha compositing modes supported by the presentation
  engine for the surface on the specified device, and at least one bit
  will be set. Opaque composition **can** be achieved in any alpha
  compositing mode by either using an image format that has no alpha
  component, or by ensuring that all pixels in the presentable images
  have an alpha value of 1.0.

- `supportedUsageFlags` is a bitmask of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) representing the
  ways the application **can** use the presentable images of a swapchain
  created with [VkPresentModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPresentModeKHR.html) set to
  `VK_PRESENT_MODE_IMMEDIATE_KHR`, `VK_PRESENT_MODE_MAILBOX_KHR`,
  `VK_PRESENT_MODE_FIFO_KHR` or `VK_PRESENT_MODE_FIFO_RELAXED_KHR` for
  the surface on the specified device.
  `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` **must** be included in the set.
  Implementations **may** support additional usages.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Supported usage flags of a presentable image when using
<code>VK_PRESENT_MODE_SHARED_DEMAND_REFRESH_KHR</code> or
<code>VK_PRESENT_MODE_SHARED_CONTINUOUS_REFRESH_KHR</code> presentation
mode are provided by <a
href="VkSharedPresentSurfaceCapabilitiesKHR.html">VkSharedPresentSurfaceCapabilitiesKHR</a>::<code>sharedPresentSupportedUsageFlags</code>.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Formulas such as min(N, <code>maxImageCount</code>) are not correct,
since <code>maxImageCount</code> <strong>may</strong> be zero.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html),
[VkCompositeAlphaFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompositeAlphaFlagsKHR.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html),
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html),
[VkSurfaceCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCapabilities2KHR.html),
[VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html),
[VkSurfaceTransformFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagsKHR.html),
[vkGetPhysicalDeviceSurfaceCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilitiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSurfaceCapabilitiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
