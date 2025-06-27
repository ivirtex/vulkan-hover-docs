# VK_SHADER_UNUSED_KHR(3) Manual Page

## Name

VK_SHADER_UNUSED_KHR - Sentinel for an unused shader index



## <a href="#_c_specification" class="anchor"></a>C Specification

`VK_SHADER_UNUSED_KHR` is a special shader index used to indicate that a
ray generation, miss, or callable shader member is not used.

``` c
#define VK_SHADER_UNUSED_KHR              (~0U)
```

or the equivalent

``` c
#define VK_SHADER_UNUSED_NV               VK_SHADER_UNUSED_KHR
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_ray_tracing_pipeline](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_ray_tracing_pipeline.html),
[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_SHADER_UNUSED_KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
