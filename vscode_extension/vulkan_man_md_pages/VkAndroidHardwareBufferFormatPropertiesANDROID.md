# VkAndroidHardwareBufferFormatPropertiesANDROID(3) Manual Page

## Name

VkAndroidHardwareBufferFormatPropertiesANDROID - Structure describing
the image format properties of an Android hardware buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To obtain format properties of an Android hardware buffer, include a
`VkAndroidHardwareBufferFormatPropertiesANDROID` structure in the
`pNext` chain of the
[VkAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferPropertiesANDROID.html)
structure passed to
[vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html).
This structure is defined as:

``` c
// Provided by VK_ANDROID_external_memory_android_hardware_buffer
typedef struct VkAndroidHardwareBufferFormatPropertiesANDROID {
    VkStructureType                  sType;
    void*                            pNext;
    VkFormat                         format;
    uint64_t                         externalFormat;
    VkFormatFeatureFlags             formatFeatures;
    VkComponentMapping               samplerYcbcrConversionComponents;
    VkSamplerYcbcrModelConversion    suggestedYcbcrModel;
    VkSamplerYcbcrRange              suggestedYcbcrRange;
    VkChromaLocation                 suggestedXChromaOffset;
    VkChromaLocation                 suggestedYChromaOffset;
} VkAndroidHardwareBufferFormatPropertiesANDROID;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `format` is the Vulkan format corresponding to the Android hardware
  buffer’s format, or `VK_FORMAT_UNDEFINED` if there is not an
  equivalent Vulkan format.

- `externalFormat` is an implementation-defined external format
  identifier for use with
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html). It **must**
  not be zero.

- `formatFeatures` describes the capabilities of this external format
  when used with an image bound to memory imported from `buffer`.

- `samplerYcbcrConversionComponents` is the component swizzle that
  **should** be used in
  [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html).

- `suggestedYcbcrModel` is a suggested color model to use in the
  [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html).

- `suggestedYcbcrRange` is a suggested numerical value range to use in
  [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html).

- `suggestedXChromaOffset` is a suggested X chroma offset to use in
  [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html).

- `suggestedYChromaOffset` is a suggested Y chroma offset to use in
  [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html).

## <a href="#_description" class="anchor"></a>Description

If the Android hardware buffer has one of the formats listed in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-android-hardware-buffer-formats"
target="_blank" rel="noopener">Format Equivalence table</a>, then
`format` **must** have the equivalent Vulkan format listed in the table.
Otherwise, `format` **may** be `VK_FORMAT_UNDEFINED`, indicating the
Android hardware buffer **can** only be used with an external format.

The `formatFeatures` member **must** include
`VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT` and at least one of
`VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT` or
`VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT`, and **should** include
`VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT` and
`VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT`.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The <code>formatFeatures</code> member only indicates the features
available when using an <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-external-android-hardware-buffer-external-formats"
target="_blank" rel="noopener">external-format image</a> created from
the Android hardware buffer. Images from Android hardware buffers with a
format other than <code>VK_FORMAT_UNDEFINED</code> are subject to the
format capabilities obtained from <a
href="vkGetPhysicalDeviceFormatProperties2.html">vkGetPhysicalDeviceFormatProperties2</a>,
and <a
href="vkGetPhysicalDeviceImageFormatProperties2.html">vkGetPhysicalDeviceImageFormatProperties2</a>
with appropriate parameters. These sets of features are independent of
each other, e.g. the external format will support sampler
Y′C<sub>B</sub>C<sub>R</sub> conversion even if the non-external format
does not, and rendering directly to the external format will not be
supported even if the non-external format does support this.</p></td>
</tr>
</tbody>
</table>

Android hardware buffers with the same external format **must** have the
same support for `VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_LINEAR_BIT`,
`VK_FORMAT_FEATURE_MIDPOINT_CHROMA_SAMPLES_BIT`,
`VK_FORMAT_FEATURE_COSITED_CHROMA_SAMPLES_BIT`,
`VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_LINEAR_FILTER_BIT`,
`VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_SEPARATE_RECONSTRUCTION_FILTER_BIT`,
and
`VK_FORMAT_FEATURE_SAMPLED_IMAGE_YCBCR_CONVERSION_CHROMA_RECONSTRUCTION_EXPLICIT_FORCEABLE_BIT`.
in `formatFeatures`. Other format features **may** differ between
Android hardware buffers that have the same external format. This allows
applications to use the same
[VkSamplerYcbcrConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversion.html) object (and
samplers and pipelines created from them) for any Android hardware
buffers that have the same external format.

If `format` is not `VK_FORMAT_UNDEFINED`, then the value of
`samplerYcbcrConversionComponents` **must** be valid when used as the
`components` member of
[VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrConversionCreateInfo.html)
with that format. If `format` is `VK_FORMAT_UNDEFINED`, all members of
`samplerYcbcrConversionComponents` **must** be the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
target="_blank" rel="noopener">identity swizzle</a>.

Implementations **may** not always be able to determine the color model,
numerical range, or chroma offsets of the image contents, so the values
in `VkAndroidHardwareBufferFormatPropertiesANDROID` are only
suggestions. Applications **should** treat these values as sensible
defaults to use in the absence of more reliable information obtained
through some other means. If the underlying physical device is also
usable via OpenGL ES with the
[`GL_OES_EGL_image_external`](https://registry.khronos.org/OpenGL/extensions/OES/OES_EGL_image_external.txt)
extension, the implementation **should** suggest values that will
produce similar sampled values as would be obtained by sampling the same
external image via `samplerExternalOES` in OpenGL ES using equivalent
sampler parameters.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Since <a
href="https://registry.khronos.org/OpenGL/extensions/OES/OES_EGL_image_external.txt"><code>GL_OES_EGL_image_external</code></a>
does not require the same sampling and conversion calculations as Vulkan
does, achieving identical results between APIs <strong>may</strong> not
be possible on some implementations.</p></td>
</tr>
</tbody>
</table>

Valid Usage (Implicit)

- <a
  href="#VUID-VkAndroidHardwareBufferFormatPropertiesANDROID-sType-sType"
  id="VUID-VkAndroidHardwareBufferFormatPropertiesANDROID-sType-sType"></a>
  VUID-VkAndroidHardwareBufferFormatPropertiesANDROID-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_FORMAT_PROPERTIES_ANDROID`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ANDROID_external_memory_android_hardware_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html),
[VkChromaLocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocation.html),
[VkComponentMapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentMapping.html),
[VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkFormatFeatureFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlags.html),
[VkSamplerYcbcrModelConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrModelConversion.html),
[VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrRange.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAndroidHardwareBufferFormatPropertiesANDROID"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
