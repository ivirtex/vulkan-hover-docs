# MTLBuffer\_id(3) Manual Page

## Name

MTLBuffer\_id - Metal MTLBuffer type reference



## [](#_c_specification)C Specification

The type `id<MTLBuffer>` is defined in Apple’s Metal framework, but to remove an unnecessary compile time dependency, an incomplete type definition of [MTLBuffer\_id](https://registry.khronos.org/vulkan/specs/latest/man/html/MTLBuffer_id.html) is provided in the Vulkan headers:

```c++
// Provided by VK_EXT_metal_objects
#ifdef __OBJC__
@protocol MTLBuffer;
typedef __unsafe_unretained id<MTLBuffer> MTLBuffer_id;
#else
typedef void* MTLBuffer_id;
#endif
```

## [](#_see_also)See Also

[VK\_EXT\_metal\_objects](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_objects.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#MTLBuffer_id).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0