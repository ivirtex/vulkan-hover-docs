# SubgroupLocalInvocationId(3) Manual Page

## Name

SubgroupLocalInvocationId - ID of the invocation within a subgroup



## <a href="#_description" class="anchor"></a>Description

`SubgroupLocalInvocationId`  
Decorating a variable with the `SubgroupLocalInvocationId` builtin
decoration will make that variable contain the index of the invocation
within the subgroup. This variable is in range \[0,`SubgroupSize`-1\].

If `VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT` is
specified, or if `module` declares SPIR-V version 1.6 or higher, and the
local workgroup size in the X dimension of the `stage` is a multiple of
<a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#interfaces-builtin-variables-sgs"
target="_blank" rel="noopener"><code>SubgroupSize</code></a>, full
subgroups are enabled for that pipeline stage. When full subgroups are
enabled, subgroups **must** be launched with all invocations active,
i.e., there is an active invocation with `SubgroupLocalInvocationId` for
each value in range \[0,`SubgroupSize`-1\].

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>There is no direct relationship between
<code>SubgroupLocalInvocationId</code> and
<code>LocalInvocationId</code> or <code>LocalInvocationIndex</code>. If
the pipeline or shader object was created with full subgroups
applications can compute their own local invocation index to serve the
same purpose:</p>
<p>index = <code>SubgroupLocalInvocationId</code> +
<code>SubgroupId</code> × <code>SubgroupSize</code></p>
<p>If full subgroups are not enabled, some subgroups may be dispatched
with inactive invocations that do not correspond to a local workgroup
invocation, making the value of index unreliable.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p><code>VK_PIPELINE_SHADER_STAGE_CREATE_REQUIRE_FULL_SUBGROUPS_BIT</code>
and <code>VK_SHADER_CREATE_REQUIRE_FULL_SUBGROUPS_BIT_EXT</code> are
effectively deprecated when compiling SPIR-V 1.6 shaders, as this
behavior is the default for Vulkan with SPIR-V 1.6. This is more aligned
with developer expectations, and avoids applications unexpectedly
breaking in the future.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a
  href="#VUID-SubgroupLocalInvocationId-SubgroupLocalInvocationId-04380"
  id="VUID-SubgroupLocalInvocationId-SubgroupLocalInvocationId-04380"></a>
  VUID-SubgroupLocalInvocationId-SubgroupLocalInvocationId-04380  
  The variable decorated with `SubgroupLocalInvocationId` **must** be
  declared using the `Input` `Storage` `Class`

- <a
  href="#VUID-SubgroupLocalInvocationId-SubgroupLocalInvocationId-04381"
  id="VUID-SubgroupLocalInvocationId-SubgroupLocalInvocationId-04381"></a>
  VUID-SubgroupLocalInvocationId-SubgroupLocalInvocationId-04381  
  The variable decorated with `SubgroupLocalInvocationId` **must** be
  declared as a scalar 32-bit integer value

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#SubgroupLocalInvocationId"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
