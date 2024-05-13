# VkSurfaceCapabilities2EXT(3) Manual Page

## Name

VkSurfaceCapabilities2EXT - Structure describing capabilities of a
surface



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSurfaceCapabilities2EXT` structure is defined as:

``` c
// Provided by VK_EXT_display_surface_counter
typedef struct VkSurfaceCapabilities2EXT {
    VkStructureType                  sType;
    void*                            pNext;
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
    VkSurfaceCounterFlagsEXT         supportedSurfaceCounters;
} VkSurfaceCapabilities2EXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

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

- `supportedSurfaceCounters` is a bitmask of
  [VkSurfaceCounterFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCounterFlagBitsEXT.html)
  indicating the supported surface counter types.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkSurfaceCapabilities2EXT-supportedSurfaceCounters-01246"
  id="VUID-VkSurfaceCapabilities2EXT-supportedSurfaceCounters-01246"></a>
  VUID-VkSurfaceCapabilities2EXT-supportedSurfaceCounters-01246  
  `supportedSurfaceCounters` **must** not include
  `VK_SURFACE_COUNTER_VBLANK_BIT_EXT` unless the surface queried is a <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#wsi-display-surfaces"
  target="_blank" rel="noopener">display surface</a>

Valid Usage (Implicit)

- <a href="#VUID-VkSurfaceCapabilities2EXT-sType-sType"
  id="VUID-VkSurfaceCapabilities2EXT-sType-sType"></a>
  VUID-VkSurfaceCapabilities2EXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SURFACE_CAPABILITIES_2_EXT`

- <a href="#VUID-VkSurfaceCapabilities2EXT-pNext-pNext"
  id="VUID-VkSurfaceCapabilities2EXT-pNext-pNext"></a>
  VUID-VkSurfaceCapabilities2EXT-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_display_surface_counter](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_display_surface_counter.html),
[VkCompositeAlphaFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCompositeAlphaFlagsKHR.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html),
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSurfaceCounterFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceCounterFlagsEXT.html),
[VkSurfaceTransformFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagBitsKHR.html),
[VkSurfaceTransformFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceTransformFlagsKHR.html),
[vkGetPhysicalDeviceSurfaceCapabilities2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceCapabilities2EXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSurfaceCapabilities2EXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
