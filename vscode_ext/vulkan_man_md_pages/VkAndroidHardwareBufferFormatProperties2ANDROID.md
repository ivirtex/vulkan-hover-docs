# VkAndroidHardwareBufferFormatProperties2ANDROID(3) Manual Page

## Name

VkAndroidHardwareBufferFormatProperties2ANDROID - Structure describing
the image format properties of an Android hardware buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

The format properties of an Android hardware buffer **can** be obtained
by including a `VkAndroidHardwareBufferFormatProperties2ANDROID`
structure in the `pNext` chain of the
[VkAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAndroidHardwareBufferPropertiesANDROID.html)
structure passed to
[vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html).
This structure is defined as:

``` c
// Provided by VK_ANDROID_external_memory_android_hardware_buffer with VK_KHR_format_feature_flags2 or VK_VERSION_1_3
typedef struct VkAndroidHardwareBufferFormatProperties2ANDROID {
    VkStructureType                  sType;
    void*                            pNext;
    VkFormat                         format;
    uint64_t                         externalFormat;
    VkFormatFeatureFlags2            formatFeatures;
    VkComponentMapping               samplerYcbcrConversionComponents;
    VkSamplerYcbcrModelConversion    suggestedYcbcrModel;
    VkSamplerYcbcrRange              suggestedYcbcrRange;
    VkChromaLocation                 suggestedXChromaOffset;
    VkChromaLocation                 suggestedYChromaOffset;
} VkAndroidHardwareBufferFormatProperties2ANDROID;
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

The bits reported in `formatFeatures` **must** include the bits reported
in the corresponding fields of
`VkAndroidHardwareBufferFormatPropertiesANDROID`::`formatFeatures`.

Valid Usage (Implicit)

- <a
  href="#VUID-VkAndroidHardwareBufferFormatProperties2ANDROID-sType-sType"
  id="VUID-VkAndroidHardwareBufferFormatProperties2ANDROID-sType-sType"></a>
  VUID-VkAndroidHardwareBufferFormatProperties2ANDROID-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_FORMAT_PROPERTIES_2_ANDROID`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_ANDROID_external_memory_android_hardware_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html),
[VK_KHR_format_feature_flags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkChromaLocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkChromaLocation.html),
[VkComponentMapping](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkComponentMapping.html),
[VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkFormatFeatureFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlags2.html),
[VkSamplerYcbcrModelConversion](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrModelConversion.html),
[VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerYcbcrRange.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAndroidHardwareBufferFormatProperties2ANDROID"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
