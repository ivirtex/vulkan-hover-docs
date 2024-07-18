# VkSamplerYcbcrModelConversion(3) Manual Page

## Name

VkSamplerYcbcrModelConversion - Color model component of a color space



## <a href="#_c_specification" class="anchor"></a>C Specification

[VkSamplerYcbcrModelConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrModelConversion.html)
defines the conversion from the source color model to the shader color
model. Possible values are:

``` c
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

``` c
// Provided by VK_KHR_sampler_ycbcr_conversion
typedef VkSamplerYcbcrModelConversion VkSamplerYcbcrModelConversionKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SAMPLER_YCBCR_MODEL_CONVERSION_RGB_IDENTITY` specifies that the
  input values to the conversion are unmodified.

- `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_IDENTITY` specifies no model
  conversion but the inputs are range expanded as for
  Y′C<sub>B</sub>C<sub>R</sub>.

- `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_709` specifies the color
  model conversion from Y′C<sub>B</sub>C<sub>R</sub> to R′G′B′ defined
  in BT.709 and described in the “BT.709 Y′C<sub>B</sub>C<sub>R</sub>
  conversion” section of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#data-format"
  target="_blank" rel="noopener">Khronos Data Format Specification</a>.

- `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_601` specifies the color
  model conversion from Y′C<sub>B</sub>C<sub>R</sub> to R′G′B′ defined
  in BT.601 and described in the “BT.601 Y′C<sub>B</sub>C<sub>R</sub>
  conversion” section of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#data-format"
  target="_blank" rel="noopener">Khronos Data Format Specification</a>.

- `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_2020` specifies the color
  model conversion from Y′C<sub>B</sub>C<sub>R</sub> to R′G′B′ defined
  in BT.2020 and described in the “BT.2020 Y′C<sub>B</sub>C<sub>R</sub>
  conversion” section of the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#data-format"
  target="_blank" rel="noopener">Khronos Data Format Specification</a>.

In the `VK_SAMPLER_YCBCR_MODEL_CONVERSION_YCBCR_*` color models, for the
input to the sampler Y′C<sub>B</sub>C<sub>R</sub> range expansion and
model conversion:

- the Y (Y′ luma) component corresponds to the G component of an RGB
  image.

- the CB (C<sub>B</sub> or “U” blue color difference) component
  corresponds to the B component of an RGB image.

- the CR (C<sub>R</sub> or “V” red color difference) component
  corresponds to the R component of an RGB image.

- the alpha component, if present, is not modified by color model
  conversion.

These rules reflect the mapping of components after the component
swizzle operation (controlled by
[VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html)::`components`).

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>For example, an “YUVA” 32-bit format comprising four 8-bit components
can be implemented as <code>VK_FORMAT_R8G8B8A8_UNORM</code> with a
component mapping:</p>
<ul>
<li><p><code>components.a</code> =
<code>VK_COMPONENT_SWIZZLE_IDENTITY</code></p></li>
<li><p><code>components.r</code> =
<code>VK_COMPONENT_SWIZZLE_B</code></p></li>
<li><p><code>components.g</code> =
<code>VK_COMPONENT_SWIZZLE_R</code></p></li>
<li><p><code>components.b</code> =
<code>VK_COMPONENT_SWIZZLE_G</code></p></li>
</ul></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkAndroidHardwareBufferFormatProperties2ANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatProperties2ANDROID.html),
[VkAndroidHardwareBufferFormatPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferFormatPropertiesANDROID.html),
[VkBufferCollectionPropertiesFUCHSIA](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCollectionPropertiesFUCHSIA.html),
[VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html),
[VkScreenBufferFormatPropertiesQNX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkScreenBufferFormatPropertiesQNX.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSamplerYcbcrModelConversion"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
