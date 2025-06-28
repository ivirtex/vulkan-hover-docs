# VkSamplerYcbcrModelConversion(3) Manual Page

## Name

VkSamplerYcbcrModelConversion - Color model component of a color space



## [](#_c_specification)C Specification

[VkSamplerYcbcrModelConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrModelConversion.html) defines the conversion from the source color model to the shader color model. Possible values are:

```c++
// Provided by VK_VERSION_1_1
typedef enum VkSamplerYcbcrModelConversion {
    VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY = 0,
    VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY = 1,
    VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709 = 2,
    VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601 = 3,
    VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020 = 4,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY_KHR = VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY_KHR = VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709_KHR = VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601_KHR = VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601,
  // Provided by VK_KHR_sampler_ycbcr_conversion
    VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020_KHR = VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020,
} VkSamplerYcbcrModelConversion;
```

or the equivalent

```c++
// Provided by VK_KHR_sampler_ycbcr_conversion
typedef VkSamplerYcbcrModelConversion VkSamplerYcbcrModelConversionKHR;
```

## [](#_description)Description

- `VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY` specifies that the input values to the conversion are unmodified.
- `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY` specifies no model conversion but the inputs are range expanded as for Y′CBCR.
- `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709` specifies the color model conversion from Y′CBCR to R′G′B′ defined in BT.709 and described in the “BT.709 Y′CBCR conversion” section of the [Khronos Data Format Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#data-format).
- `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601` specifies the color model conversion from Y′CBCR to R′G′B′ defined in BT.601 and described in the “BT.601 Y′CBCR conversion” section of the [Khronos Data Format Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#data-format).
- `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020` specifies the color model conversion from Y′CBCR to R′G′B′ defined in BT.2020 and described in the “BT.2020 Y′CBCR conversion” section of the [Khronos Data Format Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#data-format).

In the `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_*` color models, for the input to the sampler Y′CBCR range expansion and model conversion:

- the Y (Y′ luma) component corresponds to the G component of an RGB image.
- the CB (CB or “U” blue color difference) component corresponds to the B component of an RGB image.
- the CR (CR or “V” red color difference) component corresponds to the R component of an RGB image.
- the alpha component, if present, is not modified by color model conversion.

These rules reflect the mapping of components after the component swizzle operation (controlled by [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html)::`components`).

Note

For example, an “YUVA” 32-bit format comprising four 8-bit components can be implemented as `VK_FORMAT_R8G8B8A8_UNORM` with a component mapping:

- `components.a` = `VK_COMPONENT_SWIZZLE_IDENTITY`
- `components.r` = `VK_COMPONENT_SWIZZLE_B`
- `components.g` = `VK_COMPONENT_SWIZZLE_R`
- `components.b` = `VK_COMPONENT_SWIZZLE_G`

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkAndroidHardwareBufferFormatProperties2ANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatProperties2ANDROID.html), [VkAndroidHardwareBufferFormatPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferFormatPropertiesANDROID.html), [VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCollectionPropertiesFUCHSIA.html), [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html), [VkScreenBufferFormatPropertiesQNX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkScreenBufferFormatPropertiesQNX.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSamplerYcbcrModelConversion)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0