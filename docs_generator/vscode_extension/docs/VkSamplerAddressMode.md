# VkSamplerAddressMode(3) Manual Page

## Name

VkSamplerAddressMode - Specify behavior of sampling with texture coordinates outside an image



## [](#_c_specification)C Specification

Possible values of the [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html)::`addressMode*` parameters, specifying the behavior of sampling with coordinates outside the range \[0,1] for the respective u, v, or w coordinate as defined in the [Wrapping Operation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#textures-wrapping-operation) section, are:

```c++
// Provided by VK_VERSION_1_0
typedef enum VkSamplerAddressMode {
    VK_SAMPLER_ADDRESS_MODE_REPEAT = 0,
    VK_SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT = 1,
    VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE = 2,
    VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER = 3,
  // Provided by VK_VERSION_1_2, VK_KHR_sampler_mirror_clamp_to_edge
    VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE = 4,
  // Provided by VK_KHR_sampler_mirror_clamp_to_edge
  // VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE_KHR is a deprecated alias
    VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE_KHR = VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE,
} VkSamplerAddressMode;
```

## [](#_description)Description

- `VK_SAMPLER_ADDRESS_MODE_REPEAT` specifies that the repeat wrap mode will be used.
- `VK_SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT` specifies that the mirrored repeat wrap mode will be used.
- `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE` specifies that the clamp to edge wrap mode will be used.
- `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER` specifies that the clamp to border wrap mode will be used.
- `VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE` specifies that the mirror clamp to edge wrap mode will be used. This is only valid if the [`samplerMirrorClampToEdge`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-samplerMirrorClampToEdge) feature is enabled, or if the `VK_KHR_sampler_mirror_clamp_to_edge` extension is enabled.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSamplerAddressMode)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0