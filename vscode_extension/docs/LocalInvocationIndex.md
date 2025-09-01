# LocalInvocationIndex(3) Manual Page

## Name

LocalInvocationIndex - Linear local invocation index



## [](#_description)Description

`LocalInvocationIndex`

Decorating a variable with the `LocalInvocationIndex` built-in decoration will make that variable contain a one-dimensional representation of `LocalInvocationId`. This is computed as:

```c++
LocalInvocationIndex =
    LocalInvocationId.z * WorkgroupSize.x * WorkgroupSize.y +
    LocalInvocationId.y * WorkgroupSize.x +
    LocalInvocationId.x;
```

Valid Usage

- [](#VUID-LocalInvocationIndex-LocalInvocationIndex-04284)VUID-LocalInvocationIndex-LocalInvocationIndex-04284  
  The `LocalInvocationIndex` decoration **must** be used only within the `GLCompute`, `MeshEXT`, `TaskEXT`, `MeshNV`, or `TaskNV` `Execution` `Model`
- [](#VUID-LocalInvocationIndex-LocalInvocationIndex-04285)VUID-LocalInvocationIndex-LocalInvocationIndex-04285  
  The variable decorated with `LocalInvocationIndex` **must** be declared using the `Input` `Storage` `Class`
- [](#VUID-LocalInvocationIndex-LocalInvocationIndex-04286)VUID-LocalInvocationIndex-LocalInvocationIndex-04286  
  The variable decorated with `LocalInvocationIndex` **must** be declared as a scalar 32-bit integer value

## [](#_see_also)See Also

No cross-references are available

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#LocalInvocationIndex).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0