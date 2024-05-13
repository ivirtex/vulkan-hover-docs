# VkDrmFormatModifierProperties2EXT(3) Manual Page

## Name

VkDrmFormatModifierProperties2EXT - Structure specifying properties of a
format when combined with a DRM format modifier



## <a href="#_c_specification" class="anchor"></a>C Specification

The
[VkDrmFormatModifierProperties2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierProperties2EXT.html)
structure describes properties of a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) when that
format is combined with a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#glossary-drm-format-modifier"
target="_blank" rel="noopener">Linux DRM format modifier</a>. These
properties, like those of
[VkFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2.html), are independent of any
particular image.

The
[VkDrmFormatModifierPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesEXT.html)
structure is defined as:

``` c
// Provided by VK_EXT_image_drm_format_modifier with VK_KHR_format_feature_flags2 or VK_VERSION_1_3
typedef struct VkDrmFormatModifierProperties2EXT {
    uint64_t                 drmFormatModifier;
    uint32_t                 drmFormatModifierPlaneCount;
    VkFormatFeatureFlags2    drmFormatModifierTilingFeatures;
} VkDrmFormatModifierProperties2EXT;
```

## <a href="#_members" class="anchor"></a>Members

- `drmFormatModifier` is a *Linux DRM format modifier*.

- `drmFormatModifierPlaneCount` is the number of *memory planes* in any
  image created with `format` and `drmFormatModifier`. An image’s
  *memory planecount* is distinct from its *format planecount*, as
  explained below.

- `drmFormatModifierTilingFeatures` is a bitmask of
  [VkFormatFeatureFlagBits2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlagBits2.html) that are
  supported by any image created with `format` and `drmFormatModifier`.

## <a href="#_description" class="anchor"></a>Description

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_image_drm_format_modifier](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_image_drm_format_modifier.html),
[VK_KHR_format_feature_flags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_format_feature_flags2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkDrmFormatModifierPropertiesList2EXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDrmFormatModifierPropertiesList2EXT.html),
[VkFormatFeatureFlags2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatFeatureFlags2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDrmFormatModifierProperties2EXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
