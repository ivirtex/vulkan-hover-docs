# WorkgroupId(3) Manual Page

## Name

WorkgroupId - Workgroup ID of a shader



## <a href="#_description" class="anchor"></a>Description

`WorkgroupId`  
Decorating a variable with the `WorkgroupId` built-in decoration will
make that variable contain the global workgroup that the current
invocation is a member of. Each component ranges from a base value to a
base + count value, based on the parameters passed into the dispatching
commands.

Valid Usage

- <a href="#VUID-WorkgroupId-WorkgroupId-04422"
  id="VUID-WorkgroupId-WorkgroupId-04422"></a>
  VUID-WorkgroupId-WorkgroupId-04422  
  The `WorkgroupId` decoration **must** be used only within the
  `GLCompute`, `MeshEXT`, `TaskEXT`, `MeshNV`, or `TaskNV` `Execution`
  `Model`

- <a href="#VUID-WorkgroupId-WorkgroupId-04423"
  id="VUID-WorkgroupId-WorkgroupId-04423"></a>
  VUID-WorkgroupId-WorkgroupId-04423  
  The variable decorated with `WorkgroupId` **must** be declared using
  the `Input` `Storage` `Class`

- <a href="#VUID-WorkgroupId-WorkgroupId-04424"
  id="VUID-WorkgroupId-WorkgroupId-04424"></a>
  VUID-WorkgroupId-WorkgroupId-04424  
  The variable decorated with `WorkgroupId` **must** be declared as a
  three-component vector of 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#WorkgroupId"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
