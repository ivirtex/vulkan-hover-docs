# VkResolveModeFlagBits(3) Manual Page

## Name

VkResolveModeFlagBits - Bitmask indicating supported depth and stencil resolve modes



## [](#_c_specification)C Specification

Multisample values in a multisample attachment are combined according to the resolve mode used:

```c++
// Provided by VK_VERSION_1_2
typedef enum VkResolveModeFlagBits {
    VK_RESOLVE_MODE_NONE = 0,
    VK_RESOLVE_MODE_SAMPLE_ZERO_BIT = 0x00000001,
    VK_RESOLVE_MODE_AVERAGE_BIT = 0x00000002,
    VK_RESOLVE_MODE_MIN_BIT = 0x00000004,
    VK_RESOLVE_MODE_MAX_BIT = 0x00000008,
  // Provided by VK_ANDROID_external_format_resolve with VK_KHR_dynamic_rendering or VK_VERSION_1_3
    VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID = 0x00000010,
  // Provided by VK_KHR_depth_stencil_resolve
    VK_RESOLVE_MODE_NONE_KHR = VK_RESOLVE_MODE_NONE,
  // Provided by VK_KHR_depth_stencil_resolve
    VK_RESOLVE_MODE_SAMPLE_ZERO_BIT_KHR = VK_RESOLVE_MODE_SAMPLE_ZERO_BIT,
  // Provided by VK_KHR_depth_stencil_resolve
    VK_RESOLVE_MODE_AVERAGE_BIT_KHR = VK_RESOLVE_MODE_AVERAGE_BIT,
  // Provided by VK_KHR_depth_stencil_resolve
    VK_RESOLVE_MODE_MIN_BIT_KHR = VK_RESOLVE_MODE_MIN_BIT,
  // Provided by VK_KHR_depth_stencil_resolve
    VK_RESOLVE_MODE_MAX_BIT_KHR = VK_RESOLVE_MODE_MAX_BIT,
  // Provided by VK_ANDROID_external_format_resolve with VK_KHR_dynamic_rendering or VK_VERSION_1_3
  // VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID is a deprecated alias
    VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_ANDROID = VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID,
} VkResolveModeFlagBits;
```

or the equivalent

```c++
// Provided by VK_KHR_depth_stencil_resolve
typedef VkResolveModeFlagBits VkResolveModeFlagBitsKHR;
```

## [](#_description)Description

- `VK_RESOLVE_MODE_NONE` specifies that no resolve operation is done.
- `VK_RESOLVE_MODE_SAMPLE_ZERO_BIT` specifies that result of the resolve operation is equal to the value of sample 0.
- `VK_RESOLVE_MODE_AVERAGE_BIT` specifies that result of the resolve operation is the average of the sample values.
- `VK_RESOLVE_MODE_MIN_BIT` specifies that result of the resolve operation is the minimum of the sample values.
- `VK_RESOLVE_MODE_MAX_BIT` specifies that result of the resolve operation is the maximum of the sample values.
- `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID` specifies that rather than a multisample resolve, a single sampled color attachment will be downsampled into a Y′CBCR format image specified by an external Android format. Unlike other resolve modes, implementations can resolve multiple times during rendering, or even bypass writing to the color attachment altogether, as long as the final value is resolved to the resolve attachment. Values in the G, B, and R channels of the color attachment will be written to the Y, CB, and CR channels of the external format image, respectively. Chroma values are calculated as if sampling with a linear filter from the color attachment at full rate, at the location the chroma values sit according to [VkPhysicalDeviceExternalFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalFormatResolvePropertiesANDROID.html)::`externalFormatResolveChromaOffsetX`, [VkPhysicalDeviceExternalFormatResolvePropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceExternalFormatResolvePropertiesANDROID.html)::`externalFormatResolveChromaOffsetY`, and the chroma sample rate of the resolved image.

If no resolve mode is otherwise specified, `VK_RESOLVE_MODE_AVERAGE_BIT` is used.

If `VK_RESOLVE_MODE_AVERAGE_BIT` is used, and the source format is a floating-point or normalized type, the sample values for each pixel are resolved with implementation-defined numerical precision.

If the [numeric format](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#formats-numericformat) of the resolve attachment uses sRGB encoding, the implementation **should** convert samples from nonlinear to linear before averaging samples as described in the “sRGB EOTF” section of the [Khronos Data Format Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#data-format). In this case, the implementation **must** convert the linear averaged value to nonlinear before writing the resolved result to resolve attachment.

Note

No range compression or Y′CBCR model conversion is performed by `VK_RESOLVE_MODE_EXTERNAL_FORMAT_DOWNSAMPLE_BIT_ANDROID`; applications have to do these conversions themselves. Value outputs are expected to match those that would be read through a [Y′CBCR sampler using `VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-sampler-YCbCr-conversion-modelconversion). The color space that the values should be in is defined by the platform and is not exposed via Vulkan.

## [](#_see_also)See Also

[VK\_KHR\_depth\_stencil\_resolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_depth_stencil_resolve.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkRenderingAttachmentInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAttachmentInfo.html), [VkResolveModeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkResolveModeFlags.html), [VkSubpassDescriptionDepthStencilResolve](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescriptionDepthStencilResolve.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkResolveModeFlagBits)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0