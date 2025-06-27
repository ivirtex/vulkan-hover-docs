# VkDisplayPlaneCapabilitiesKHR(3) Manual Page

## Name

VkDisplayPlaneCapabilitiesKHR - Structure describing capabilities of a
mode and plane combination



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDisplayPlaneCapabilitiesKHR` structure is defined as:

``` c
// Provided by VK_KHR_display
typedef struct VkDisplayPlaneCapabilitiesKHR {
    VkDisplayPlaneAlphaFlagsKHR    supportedAlpha;
    VkOffset2D                     minSrcPosition;
    VkOffset2D                     maxSrcPosition;
    VkExtent2D                     minSrcExtent;
    VkExtent2D                     maxSrcExtent;
    VkOffset2D                     minDstPosition;
    VkOffset2D                     maxDstPosition;
    VkExtent2D                     minDstExtent;
    VkExtent2D                     maxDstExtent;
} VkDisplayPlaneCapabilitiesKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `supportedAlpha` is a bitmask of
  [VkDisplayPlaneAlphaFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneAlphaFlagBitsKHR.html)
  describing the supported alpha blending modes.

- `minSrcPosition` is the minimum source rectangle offset supported by
  this plane using the specified mode.

- `maxSrcPosition` is the maximum source rectangle offset supported by
  this plane using the specified mode. The `x` and `y` components of
  `maxSrcPosition` **must** each be greater than or equal to the `x` and
  `y` components of `minSrcPosition`, respectively.

- `minSrcExtent` is the minimum source rectangle size supported by this
  plane using the specified mode.

- `maxSrcExtent` is the maximum source rectangle size supported by this
  plane using the specified mode.

- `minDstPosition`, `maxDstPosition`, `minDstExtent`, `maxDstExtent` all
  have similar semantics to their corresponding `*Src*` equivalents, but
  apply to the output region within the mode rather than the input
  region within the source image. Unlike the `*Src*` offsets,
  `minDstPosition` and `maxDstPosition` **may** contain negative values.

## <a href="#_description" class="anchor"></a>Description

The minimum and maximum position and extent fields describe the
implementation limits, if any, as they apply to the specified display
mode and plane. Vendors **may** support displaying a subset of a
swapchain’s presentable images on the specified display plane. This is
expressed by returning `minSrcPosition`, `maxSrcPosition`,
`minSrcExtent`, and `maxSrcExtent` values that indicate a range of
possible positions and sizes which **may** be used to specify the region
within the presentable images that source pixels will be read from when
creating a swapchain on the specified display mode and plane.

Vendors **may** also support mapping the presentable images’ content to
a subset or superset of the visible region in the specified display
mode. This is expressed by returning `minDstPosition`, `maxDstPosition`,
`minDstExtent` and `maxDstExtent` values that indicate a range of
possible positions and sizes which **may** be used to describe the
region within the display mode that the source pixels will be mapped to.

Other vendors **may** support only a 1-1 mapping between pixels in the
presentable images and the display mode. This **may** be indicated by
returning (0,0) for `minSrcPosition`, `maxSrcPosition`,
`minDstPosition`, and `maxDstPosition`, and (display mode width, display
mode height) for `minSrcExtent`, `maxSrcExtent`, `minDstExtent`, and
`maxDstExtent`.

The value `supportedAlpha` **must** contain at least one valid
[VkDisplayPlaneAlphaFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneAlphaFlagBitsKHR.html)
bit.

These values indicate the limits of the implementation’s individual
fields. Not all combinations of values within the offset and extent
ranges returned in `VkDisplayPlaneCapabilitiesKHR` are guaranteed to be
supported. Presentation requests specifying unsupported combinations
**may** fail.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html),
[VkDisplayPlaneAlphaFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneAlphaFlagsKHR.html),
[VkDisplayPlaneCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneCapabilities2KHR.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkOffset2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOffset2D.html),
[vkGetDisplayPlaneCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDisplayPlaneCapabilitiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDisplayPlaneCapabilitiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
