# VkSamplerAddressMode(3) Manual Page

## Name

VkSamplerAddressMode - Specify behavior of sampling with texture
coordinates outside an image



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of the
[VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)::`addressMode*`
parameters, specifying the behavior of sampling with coordinates outside
the range \[0,1\] for the respective u, v, or w coordinate as defined in
the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-wrapping-operation"
target="_blank" rel="noopener">Wrapping Operation</a> section, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkSamplerAddressMode {
    VK_SAMPLER_ADDRESS_MODE_REPEAT = 0,
    VK_SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT = 1,
    VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE = 2,
    VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER = 3,
  // Provided by VK_VERSION_1_2, VK_KHR_sampler_mirror_clamp_to_edge
    VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE = 4,
  // Provided by VK_KHR_sampler_mirror_clamp_to_edge
    VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE_KHR = VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE,
} VkSamplerAddressMode;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SAMPLER_ADDRESS_MODE_REPEAT` specifies that the repeat wrap mode
  will be used.

- `VK_SAMPLER_ADDRESS_MODE_MIRRORED_REPEAT` specifies that the mirrored
  repeat wrap mode will be used.

- `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_EDGE` specifies that the clamp to
  edge wrap mode will be used.

- `VK_SAMPLER_ADDRESS_MODE_CLAMP_TO_BORDER` specifies that the clamp to
  border wrap mode will be used.

- `VK_SAMPLER_ADDRESS_MODE_MIRROR_CLAMP_TO_EDGE` specifies that the
  mirror clamp to edge wrap mode will be used. This is only valid if <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-samplerMirrorClampToEdge"
  target="_blank" rel="noopener"><code>samplerMirrorClampToEdge</code></a>
  is enabled, or if the
  [`VK_KHR_sampler_mirror_clamp_to_edge`](VK_KHR_sampler_mirror_clamp_to_edge.html)
  extension is enabled.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSamplerAddressMode"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
