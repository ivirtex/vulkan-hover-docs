# LocalInvocationIndex(3) Manual Page

## Name

LocalInvocationIndex - Linear local invocation index



## <a href="#_description" class="anchor"></a>Description

`LocalInvocationIndex`  
Decorating a variable with the `LocalInvocationIndex` built-in
decoration will make that variable contain a one-dimensional
representation of `LocalInvocationId`. This is computed as:

``` c
LocalInvocationIndex =
    LocalInvocationId.z * WorkgroupSize.x * WorkgroupSize.y +
    LocalInvocationId.y * WorkgroupSize.x +
    LocalInvocationId.x;
```

Valid Usage

- <a href="#VUID-LocalInvocationIndex-LocalInvocationIndex-04284"
  id="VUID-LocalInvocationIndex-LocalInvocationIndex-04284"></a>
  VUID-LocalInvocationIndex-LocalInvocationIndex-04284  
  The `LocalInvocationIndex` decoration **must** be used only within the
  `GLCompute`, `MeshEXT`, `TaskEXT`, `MeshNV`, or `TaskNV` `Execution`
  `Model`

- <a href="#VUID-LocalInvocationIndex-LocalInvocationIndex-04285"
  id="VUID-LocalInvocationIndex-LocalInvocationIndex-04285"></a>
  VUID-LocalInvocationIndex-LocalInvocationIndex-04285  
  The variable decorated with `LocalInvocationIndex` **must** be
  declared using the `Input` `Storage` `Class`

- <a href="#VUID-LocalInvocationIndex-LocalInvocationIndex-04286"
  id="VUID-LocalInvocationIndex-LocalInvocationIndex-04286"></a>
  VUID-LocalInvocationIndex-LocalInvocationIndex-04286  
  The variable decorated with `LocalInvocationIndex` **must** be
  declared as a scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#LocalInvocationIndex"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
