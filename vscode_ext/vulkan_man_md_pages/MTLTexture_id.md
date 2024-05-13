# MTLTexture_id(3) Manual Page

## Name

MTLTexture_id - Metal MTLTexture type reference



## <a href="#_c_specification" class="anchor"></a>C Specification

The type `id<MTLTexture>` is defined in Apple’s Metal framework, but to
remove an unnecessary compile time dependency, an incomplete type
definition of [MTLTexture_id](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/MTLTexture_id.html) is provided in the
Vulkan headers:

``` c
// Provided by VK_EXT_metal_objects
#ifdef __OBJC__
@protocol MTLTexture;
typedef __unsafe_unretained id<MTLTexture> MTLTexture_id;
#else
typedef void* MTLTexture_id;
#endif
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_objects](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_objects.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#MTLTexture_id"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
