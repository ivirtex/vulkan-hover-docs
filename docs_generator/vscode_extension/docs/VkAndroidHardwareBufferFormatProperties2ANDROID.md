# VkAndroidHardwareBufferFormatProperties2ANDROID(3) Manual Page

## Name

VkAndroidHardwareBufferFormatProperties2ANDROID - Structure describing the image format properties of an Android hardware buffer



## [](#_c_specification)C Specification

The format properties of an Android hardware buffer **can** be obtained by including a `VkAndroidHardwareBufferFormatProperties2ANDROID` structure in the `pNext` chain of the [VkAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAndroidHardwareBufferPropertiesANDROID.html) structure passed to [vkGetAndroidHardwareBufferPropertiesANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAndroidHardwareBufferPropertiesANDROID.html). This structure is defined as:

```c++
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

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `format` is the Vulkan format corresponding to the Android hardware buffer’s format, or `VK_FORMAT_UNDEFINED` if there is not an equivalent Vulkan format.
- `externalFormat` is an implementation-defined external format identifier for use with [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html). It **must** not be zero.
- `formatFeatures` describes the capabilities of this external format when used with an image bound to memory imported from `buffer`.
- `samplerYcbcrConversionComponents` is the component swizzle that **should** be used in [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html).
- `suggestedYcbcrModel` is a suggested color model to use in the [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html).
- `suggestedYcbcrRange` is a suggested numerical value range to use in [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html).
- `suggestedXChromaOffset` is a suggested X chroma offset to use in [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html).
- `suggestedYChromaOffset` is a suggested Y chroma offset to use in [VkSamplerYcbcrConversionCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrConversionCreateInfo.html).

## [](#_description)Description

The bits reported in `formatFeatures` **must** include the bits reported in the corresponding fields of `VkAndroidHardwareBufferFormatPropertiesANDROID`::`formatFeatures`.

Valid Usage (Implicit)

- [](#VUID-VkAndroidHardwareBufferFormatProperties2ANDROID-sType-sType)VUID-VkAndroidHardwareBufferFormatProperties2ANDROID-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ANDROID_HARDWARE_BUFFER_FORMAT_PROPERTIES_2_ANDROID`

## [](#_see_also)See Also

[VK\_ANDROID\_external\_memory\_android\_hardware\_buffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ANDROID_external_memory_android_hardware_buffer.html), [VK\_KHR\_format\_feature\_flags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_format_feature_flags2.html), [VK\_VERSION\_1\_3](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_3.html), [VkChromaLocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VkChromaLocation.html), [VkComponentMapping](https://registry.khronos.org/vulkan/specs/latest/man/html/VkComponentMapping.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkFormatFeatureFlags2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormatFeatureFlags2.html), [VkSamplerYcbcrModelConversion](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrModelConversion.html), [VkSamplerYcbcrRange](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerYcbcrRange.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAndroidHardwareBufferFormatProperties2ANDROID)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0