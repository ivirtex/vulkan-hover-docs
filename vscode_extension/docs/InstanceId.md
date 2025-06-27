# InstanceId(3) Manual Page

## Name

InstanceId - Id associated with an intersected instance



## <a href="#_description" class="anchor"></a>Description

`InstanceId`  
Decorating a variable in an intersection, any-hit, or closest hit shader
with the `InstanceId` decoration will make that variable contain the
index of the instance that intersects the current ray.

Valid Usage

- <a href="#VUID-InstanceId-InstanceId-04254"
  id="VUID-InstanceId-InstanceId-04254"></a>
  VUID-InstanceId-InstanceId-04254  
  The `InstanceId` decoration **must** be used only within the
  `IntersectionKHR`, `AnyHitKHR`, or `ClosestHitKHR` `Execution` `Model`

- <a href="#VUID-InstanceId-InstanceId-04255"
  id="VUID-InstanceId-InstanceId-04255"></a>
  VUID-InstanceId-InstanceId-04255  
  The variable decorated with `InstanceId` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-InstanceId-InstanceId-04256"
  id="VUID-InstanceId-InstanceId-04256"></a>
  VUID-InstanceId-InstanceId-04256  
  The variable decorated with `InstanceId` **must** be declared as a
  scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#InstanceId"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
