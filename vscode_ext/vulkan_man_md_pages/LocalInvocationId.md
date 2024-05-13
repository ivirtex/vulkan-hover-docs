# LocalInvocationId(3) Manual Page

## Name

LocalInvocationId - Local invocation ID



## <a href="#_description" class="anchor"></a>Description

`LocalInvocationId`  
Decorating a variable with the `LocalInvocationId` built-in decoration
will make that variable contain the location of the current cluster
culling, task, mesh, or compute shader invocation within the local
workgroup. Each component ranges from zero through to the size of the
workgroup in that dimension minus one.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>If the size of the workgroup in a particular dimension is one, then
the <code>LocalInvocationId</code> in that dimension will be zero. If
the workgroup is effectively two-dimensional, then
<code>LocalInvocationId.z</code> will be zero. If the workgroup is
effectively one-dimensional, then both <code>LocalInvocationId.y</code>
and <code>LocalInvocationId.z</code> will be zero.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-LocalInvocationId-LocalInvocationId-04281"
  id="VUID-LocalInvocationId-LocalInvocationId-04281"></a>
  VUID-LocalInvocationId-LocalInvocationId-04281  
  The `LocalInvocationId` decoration **must** be used only within the
  `GLCompute`, `MeshEXT`, `TaskEXT`, `MeshNV`, or `TaskNV` `Execution`
  `Model`

- <a href="#VUID-LocalInvocationId-LocalInvocationId-04282"
  id="VUID-LocalInvocationId-LocalInvocationId-04282"></a>
  VUID-LocalInvocationId-LocalInvocationId-04282  
  The variable decorated with `LocalInvocationId` **must** be declared
  using the `Input` `Storage` `Class`

- <a href="#VUID-LocalInvocationId-LocalInvocationId-04283"
  id="VUID-LocalInvocationId-LocalInvocationId-04283"></a>
  VUID-LocalInvocationId-LocalInvocationId-04283  
  The variable decorated with `LocalInvocationId` **must** be declared
  as a three-component vector of 32-bit integer values

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#LocalInvocationId"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
