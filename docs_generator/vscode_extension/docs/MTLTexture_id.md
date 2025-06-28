# MTLTexture\_id(3) Manual Page

## Name

MTLTexture\_id - Metal MTLTexture type reference



## [](#_c_specification)C Specification

The type `id<MTLTexture>` is defined in Apple’s Metal framework, but to remove an unnecessary compile time dependency, an incomplete type definition of [MTLTexture\_id](https://registry.khronos.org/vulkan/specs/latest/man/html/MTLTexture_id.html) is provided in the Vulkan headers:

```c++
// Provided by VK_EXT_metal_objects
#ifdef __OBJC__
@protocol MTLTexture;
typedef __unsafe_unretained id<MTLTexture> MTLTexture_id;
#else
typedef void* MTLTexture_id;
#endif
```

## [](#_see_also)See Also

[VK\_EXT\_metal\_objects](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_objects.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#MTLTexture_id)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0