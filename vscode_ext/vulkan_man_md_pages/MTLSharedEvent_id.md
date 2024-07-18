# MTLSharedEvent_id(3) Manual Page

## Name

MTLSharedEvent_id - Metal MTLSharedEvent type reference



## <a href="#_c_specification" class="anchor"></a>C Specification

The type `id<MTLSharedEvent>` is defined in Apple’s Metal framework, but
to remove an unnecessary compile time dependency, an incomplete type
definition of [MTLSharedEvent_id](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/MTLSharedEvent_id.html) is provided in
the Vulkan headers:

``` c
// Provided by VK_EXT_metal_objects
#ifdef __OBJC__
@protocol MTLSharedEvent;
typedef __unsafe_unretained id<MTLSharedEvent> MTLSharedEvent_id;
#else
typedef void* MTLSharedEvent_id;
#endif
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_objects](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_objects.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#MTLSharedEvent_id"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
