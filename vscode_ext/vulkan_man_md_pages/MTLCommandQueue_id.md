# MTLCommandQueue_id(3) Manual Page

## Name

MTLCommandQueue_id - Metal MTLCommandQueue type reference



## <a href="#_c_specification" class="anchor"></a>C Specification

The type `id<MTLCommandQueue>` is defined in Apple’s Metal framework,
but to remove an unnecessary compile time dependency, an incomplete type
definition of [MTLCommandQueue_id](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/MTLCommandQueue_id.html) is provided
in the Vulkan headers:

``` c
// Provided by VK_EXT_metal_objects
#ifdef __OBJC__
@protocol MTLCommandQueue;
typedef __unsafe_unretained id<MTLCommandQueue> MTLCommandQueue_id;
#else
typedef void* MTLCommandQueue_id;
#endif
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_objects](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_objects.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#MTLCommandQueue_id"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
